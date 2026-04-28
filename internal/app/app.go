package app

import (
	"context"
	"log/slog"

	"github.com/leevroko/sql_gen/internal/config"
	"github.com/leevroko/sql_gen/internal/generators"
	"github.com/leevroko/sql_gen/internal/lib/helpers/sl"
	"github.com/leevroko/sql_gen/internal/logger"
	"github.com/leevroko/sql_gen/internal/storage/postgres"
)

type DBInterface interface {
	GetEntryCount(tableName string) (int, error)
}

type App struct {
	context context.Context

	db			DBInterface
	schema 		config.DbSchema
	generator 	*generators.DataGenerationController

	logger 	*slog.Logger
}

func NewApp(cfg config.Config) *App {
	defer slog.Default().Info("App started")
	app := &App{
		context: context.TODO(),
		logger: logger.NewLogger(cfg.LogLevel),
		schema: cfg.Schema,
	}

	db := postgres.New(
		app.context, 
		cfg.DbHost, 
		cfg.DbPort, 
		cfg.DbUsername, 
		cfg.DbPassword, 
		cfg.DbName, 
		cfg.Schema,
		app.logger.WithGroup("postgres"),
	)

	app.db = db 	
	app.generator = generators.NewDataGenerationcontroller(db, cfg.AppConfig.MaxEntries, app.logger.WithGroup("generator"))
	
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
			a.generator.Populate(a.context, table, table.EntryCount - count)
		}
	}
}
