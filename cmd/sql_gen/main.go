package main

import (
	"github.com/leevroko/sql_gen/internal/config"

	applib "github.com/leevroko/sql_gen/internal/app"
)

func main() {
	cfg := config.ParseConfigFromOsArgs()
	app := applib.NewApp(cfg)
	
	app.Run()
}
