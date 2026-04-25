package main

import (
	"github.com/leevroko/sql_gen/internal/config"
	"github.com/leevroko/sql_gen/internal/storage/postgres"

	applib "github.com/leevroko/sql_gen/internal/app"
)

func main() {
	cfg := config.ParseConfigFromOsArgs()
	app := applib.NewApp(cfg)
	
	conn := postgres.New(
		app.Context, 
		cfg.DbHost, 
		cfg.DbPort, 
		cfg.DbUsername, 
		cfg.DbPassword, 
		cfg.DbName, 
		cfg.Schema,
		app.GiveLogger("postgres"),
	)
	_ = conn
	
	app.Run()
}
