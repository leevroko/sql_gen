package app

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/leevroko/sql_gen/internal/callbyname"
	"github.com/leevroko/sql_gen/internal/config"
	"github.com/leevroko/sql_gen/internal/generators"
	"github.com/leevroko/sql_gen/internal/lib/helpers/sl"
	"github.com/leevroko/sql_gen/internal/logger"
	"github.com/leevroko/sql_gen/internal/storage/postgres"
)

const (
	EntriesLimit = 10_000
)

type App struct {
	context context.Context
	db		*postgres.PostgresDb
	
	schema 	config.DbSchema

	logger 	*slog.Logger
}

func NewApp(cfg config.Config) *App {
	defer slog.Default().Info("App started")
	app := &App{
		context: context.TODO(),
		logger: logger.NewLogger(cfg.LogLevel),
		schema: cfg.Schema,
	}

	app.db = postgres.New(
		app.context, 
		cfg.DbHost, 
		cfg.DbPort, 
		cfg.DbUsername, 
		cfg.DbPassword, 
		cfg.DbName, 
		cfg.Schema,
		app.GiveLogger("postgres"),
	)

	return app
}

func (a *App) Run() {
	for _, table := range a.schema.Tables {
		count, err := a.db.GetEntryCount(table.Name)
		if err != nil {
			a.logger.Error("could not get entry count for a table", slog.String("table", table.Name), sl.Error(err))
		}
		a.logger.Debug("got entry count for a table", slog.String("table", table.Name), slog.Int("count", count))

		if count < table.EntryCount {
			a.populate(table, table.EntryCount - count)
		}
	}
}

func (a *App) GiveLogger(source string) *slog.Logger {
	return a.logger.With(slog.String("source", source))
}

func (a *App) populate(tableSchema config.TableSchema, entriesAsked int) error {
	if entriesAsked > EntriesLimit {
		panic(fmt.Sprintf("sql_gen currently does not support generating more than %v values - asked for %v", EntriesLimit, entriesAsked))
	}

	stmtbuilder := strings.Builder{} 
	stmtbuilder.WriteString("INSERT INTO " + tableSchema.Name + "\n")
	stmtbuilder.WriteString("VALUES\n")

	atLeastOneAdded := false
	for _ = range entriesAsked {
		if atLeastOneAdded {
			stmtbuilder.WriteString(",\n\t(")
		} else {
			stmtbuilder.WriteString("\t(")
		}
		atLeastOneFieldProcessed := false
		for _, column := range tableSchema.Fields {
			values := callbyname.CallByName(generators.Functions, column.GeneratedType)
			if atLeastOneFieldProcessed {
				if column.Type == "text" {
					stmtbuilder.WriteString(fmt.Sprintf(", '%v'", values[0]))
				} else {
					stmtbuilder.WriteString(fmt.Sprintf(", %v", values[0]))
				}
			} else {
				if column.Type == "text" {
					stmtbuilder.WriteString(fmt.Sprintf("'%v'", values[0]))
				} else {
					stmtbuilder.WriteString(fmt.Sprintf("%v", values[0]))
				}
			}
			atLeastOneFieldProcessed = true
		}
		stmtbuilder.WriteString(")")
		atLeastOneAdded = true
	}
	// stmtbuilder.WriteString(";")
	
	stmt := stmtbuilder.String()
	cmdTags, err := a.db.ConnPool.Exec(a.context, stmt)
	if err != nil {
		a.logger.Error("Could not insert values", sl.Error(err), slog.String("sql", stmt))
		return err
	}
	a.logger.Info("populated succesfully", slog.Any("cmdTags", cmdTags), slog.String("table", tableSchema.Name))
	return nil
}
