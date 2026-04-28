package generators

import (
	"log/slog"
	"context"
	"strings"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/leevroko/sql_gen/internal/config"
	"github.com/leevroko/sql_gen/internal/callbyname"
	"github.com/leevroko/sql_gen/internal/lib/helpers/sl"
)

type QueryExecutor interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

type DataGenerationController struct {
	db 			QueryExecutor
	entryLimit 	int
	logger 		*slog.Logger
}

func NewDataGenerationcontroller(db QueryExecutor, entryLimit int, logger *slog.Logger) *DataGenerationController {
	return &DataGenerationController{
		db: db, 
		entryLimit: entryLimit,
		logger: logger,
	}
}

func (d *DataGenerationController) Populate(ctx context.Context, tableSchema config.TableSchema, entriesAsked int) error {
	if entriesAsked > d.entryLimit {
		panic(fmt.Sprintf("sql_gen currently does not support generating more than %v values - asked for %v", d.entryLimit, entriesAsked))
	}

	stmts := d.createQueries(tableSchema, entriesAsked)

	for _, stmt := range stmts {
		cmdTags, err := d.db.Exec(ctx, stmt)
		if err != nil {
			d.logger.Error("Could not insert values", sl.Error(err), slog.String("sql", stmt))
			return err
		}
		d.logger.Info("populated succesfully", slog.Any("cmdTags", cmdTags), slog.String("table", tableSchema.Name))
	}
	return nil
}

func (d *DataGenerationController) createQueries(tableSchema config.TableSchema, entriesAsked int) []string {
	stmtbuilder := strings.Builder{} 
	stmtbuilder.WriteString("INSERT INTO " + tableSchema.Name + "\n")
	stmtbuilder.WriteString("VALUES\n")

	atLeastOneAdded := false
	for range entriesAsked {
		if atLeastOneAdded {
			stmtbuilder.WriteString(",\n\t(")
		} else {
			stmtbuilder.WriteString("\t(")
		}
		atLeastOneFieldProcessed := false
		for _, column := range tableSchema.Fields {
			values := callbyname.CallByName(GeneratorFunctions, column.GeneratedType)
			if atLeastOneFieldProcessed {
				if column.Type == "text" {
					fmt.Fprintf(&stmtbuilder, ", '%v'", values[0])
				} else {
					fmt.Fprintf(&stmtbuilder, ", %v", values[0])
				}
			} else {
				if column.Type == "text" {
					fmt.Fprintf(&stmtbuilder, "'%v'", values[0])
				} else {
					fmt.Fprintf(&stmtbuilder, "%v", values[0])
				}
			}
			atLeastOneFieldProcessed = true
		}
		stmtbuilder.WriteString(")")
		atLeastOneAdded = true
	}
	
	stmt := stmtbuilder.String()
	stmts := []string{stmt}
	return stmts
}
