package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/leevroko/sql_gen/internal/config"
)

type PostgresDb struct {
	ConnPool	*pgxpool.Pool
	cfg 		config.DbSchema
	ctx 		context.Context
	logger 		*slog.Logger
}

func New(ctx context.Context, host, port, username, password, database string, cfg config.DbSchema, logger *slog.Logger) *PostgresDb {
	hostPort := host + ":" + port
	connString := fmt.Sprintf("postgres://%v:%v@%v/%v", username, password, hostPort, database)
	
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		panic(err.Error())
	}
	
	db := &PostgresDb{
		ctx: ctx,
		ConnPool: pool,
		cfg: cfg,
		logger: logger,
	}

	schemaErr := db.validateSchema(cfg)
	if schemaErr != nil {
		panic(schemaErr.Error())
	}

	return db
}

func (p *PostgresDb) TableExists(tableName string) (bool, error) {
	stmt := `
	SELECT EXISTS (
		SELECT 1 FROM information_schema.tables
		WHERE table_schema = $1 and table_name = $2
	)`

	tableFound := false
	err := p.ConnPool.QueryRow(p.ctx, stmt, "public", tableName).Scan(&tableFound)
	if err != nil {
		return false, err
	}
	return tableFound, nil
}

func (p *PostgresDb) validateSchema(schema config.DbSchema) error {
	p.logger.Debug("Got DbSchema", slog.Any("schema", schema))
	for _, table := range schema.Tables {
		exists, err := p.TableExists(table.Name)
		if err != nil {
			panic(fmt.Errorf("Something went wrong during checking the existance of a table in the schema validation procedure: %w", err).Error())
		}

		p.logger.Debug("Examining table existance", slog.String("table", table.Name), slog.Bool("exists", exists))
		if !exists {
			stmtBuilder := strings.Builder{}
			stmtBuilder.WriteString("CREATE TABLE ")
			stmtBuilder.WriteString(table.Name)
			stmtBuilder.WriteString(" (\n")
			doneAtLeastOnce := false
			for _, field := range table.Fields {
				if !validateField(field) {
					panic(fmt.Errorf("Field %v did not pass validation", field.Name))
				}

				if !doneAtLeastOnce {
					stmtBuilder.WriteString("\t" + field.Name + " " + field.Type)
				} else {
					stmtBuilder.WriteString(",\n\t" + field.Name + " " + field.Type)
				}
				if field.Constraint != "" {
					stmtBuilder.WriteString(" " + field.Constraint)
				}
				
				doneAtLeastOnce = true
			}
			stmtBuilder.WriteString("\n);")
			if !doneAtLeastOnce {
				panic(fmt.Sprintf("table %v does not have any columns specified", table.Name))
			}
			stmt := stmtBuilder.String()
			p.logger.Debug("Table creation statement ready", slog.String("sql", stmt))
			commandTag, err := p.ConnPool.Exec(p.ctx, stmt)
			if err != nil {
				panic(fmt.Errorf("error during validation: %w", err).Error())
			}
			p.logger.Debug("Created table", slog.Any("commandTag", commandTag))
		}
	}
	return nil
}

func (p *PostgresDb) GetEntryCount(tableName string) (int, error) {
	stmt := fmt.Sprintf("SELECT COUNT(*) FROM %v;", tableName)
	
	var count int
	err := p.ConnPool.QueryRow(p.ctx, stmt).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func validateField(field config.ColumnSchema) bool {
	return ValidateSimpleType(field.Type)
}
