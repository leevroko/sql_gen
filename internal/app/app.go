package app

import (
	"context"
	"log/slog"

	"github.com/leevroko/sql_gen/internal/config"
	"github.com/leevroko/sql_gen/internal/logger"
)

type App struct {
	Context context.Context
	
	logger 	*slog.Logger
}

func NewApp(cfg config.Config) *App {
	defer slog.Default().Info("App started")
	return &App{
		Context: context.TODO(),
		logger: logger.NewLogger(cfg.LogLevel),
	}
}

func (a *App) Run() {
	slog.Error("Not implemented yet")	
}

func (a *App) GiveLogger(source string) *slog.Logger {
	return a.logger.With(slog.String("source", source))
}
