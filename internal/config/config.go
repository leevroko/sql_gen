package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/ilyakaznacheev/cleanenv"

	"github.com/leevroko/sql_gen/internal/lib/helpers/sl"
)

type Config struct {
	LogLevel 	slog.Level
	Schema 		DbSchema
	
	ConnConfig
}

type ConnConfig struct {
	DbHost 		string
	DbUsername 	string 
	DbPassword 	string
	DbPort 		string
	DbName 		string
}

type ColumnSchema struct {
	Name 			string	`yaml:"name" json:"name" env-required:"true"`
	Type 			string	`yaml:"type" json:"type" env-required:"true"`
	GeneratedType	string	`yaml:"generated_type" json:"generated_type" env-required:"true"`
	Constraint 		string	`yaml:"constraint,omitempty" json:"constraint,omitempty"`
}

type TableSchema struct {
	Name 		string 			`yaml:"name" json:"name" env-required:"true"`
	Fields 		[]ColumnSchema	`yaml:"fields" json:"fields" env-required:"true"`
	EntryCount	int				`yaml:"require_entry_count" json:"require_entry_count" env-required:"true"`
}

type DbSchema struct {
	Name 	string			`yaml:"name" json:"name" env-required:"true"`
	Tables 	[]TableSchema	`yaml:"tables" json:"tables" env-required:"true"`
}

const (
	helpMessage = `
Usage: 
	sql_gen [flags] config_file
Flags:
	-h or --help - print this message
	-l or --log-level [level] - Debug, Info, Warn or Error
	-d or --db-host [host] - host address
	-u or --db-user [username] - db user username
	-p or --db-port [port] - db port
	-n or --db-name [name] - db name
	-e or --db-password-env [env variable name] - a name of the env variable that stores the db password for the given user
	`
)

const (
	argNoFlag = iota
	argFlagLogLevel
	argFlagDbHost
	argFlagDbUser
	argFlagDbPort
	argFlagDbPassword
	argFlagDbName
)

func ParseConfigFromOsArgs() Config {
	config := Config{}
	args := os.Args
	currentFlag := argNoFlag
	for i, arg := range args {
		if len(args) == 1 {
			exitWithError(errors.New("no config path given"))
		}

		if i == 0 {
			continue
		}

		if arg == "--help" || arg == "-h" {
			slog.Default().Info(helpMessage)
			os.Exit(0)
		}

		if arg == "-l" || arg == "--log-level" {
			currentFlag = argFlagLogLevel
			continue
		}  
		if currentFlag == argFlagLogLevel {
			if 	arg == "Debug" || arg == "Info" || arg == "Warn" || arg == "Error" {
				level, err := getLogLevel(arg)	
				if err != nil {
					exitWithError(err)
				}
				config.LogLevel = level
			} else {
				exitWithError(errors.New("invalid log level"))
			}
			currentFlag = argNoFlag
			continue
		}

		if arg == "-d" || arg == "--db-host" {
			currentFlag = argFlagDbHost
			continue
		}  
		if currentFlag == argFlagDbHost {
			config.DbHost = arg
			currentFlag = argNoFlag
			continue
		}

		if arg == "-u" || arg == "--db-user" {
			currentFlag = argFlagDbUser
			continue
		}  
		if currentFlag == argFlagDbUser {
			config.DbUsername = arg
			currentFlag = argNoFlag
			continue
		}

		if arg == "-p" || arg == "--db-port" {
			currentFlag = argFlagDbPort
			continue
		}  
		if currentFlag == argFlagDbPort {
			config.DbPort = arg
			currentFlag = argNoFlag
			continue
		}

		if arg == "-n" || arg == "--db-name" {
			currentFlag = argFlagDbName
			continue
		}  
		if currentFlag == argFlagDbName {
			config.DbName = arg
			currentFlag = argNoFlag
			continue
		}

		if arg == "-e" || arg == "--db-password-env" {
			currentFlag = argFlagDbPassword
			continue
		}  
		if currentFlag == argFlagDbPassword {
			pw := os.Getenv(arg)
			config.DbPassword = pw
			currentFlag = argNoFlag
			continue
		}

		if i == (len(args) - 1) {
			if arg == "" {
				exitWithError(errors.New("config_file was not given"))
			}

			if _, err := os.Stat(arg); os.IsNotExist(err) {
				exitWithError(fmt.Errorf("config file does not exist: %v", err.Error()))
			}

			if err := cleanenv.ReadConfig(arg, &config.Schema); err != nil {
				exitWithError(fmt.Errorf("cannot read config: %v", err.Error()))
			}
			slog.Default().Info("config parsed successfully")
			continue
		}

		exitWithError(fmt.Errorf("unknkown flag: %v", arg))
	}

	return config
}

func exitWithError(err error) {
	slog.Default().Error("Error during args parsing", sl.Error(err))
	slog.Default().Error(helpMessage)
	os.Exit(1)
}

func getLogLevel(arg string) (slog.Level, error) {
	switch arg {
	case "Debug":
		return slog.LevelDebug, nil
	case "Info":
		return slog.LevelInfo, nil
	case "Warn":
		return slog.LevelWarn, nil
	case "Error":
		return slog.LevelError, nil
	default: 
		return slog.LevelDebug, errors.New("wrong log level")
	}
}
