package configs

import (
	"os"
	"strconv"
)

type cfg struct {
	Env   string
	Local bool
}

type (
	appCfg struct {
		Host    string
		Port    string
		Env     string
		Timeout uint
		Prefix  string
		Debug   bool
	}

	SystemCfg struct {
		PathSaveFile   string
		PathRootAssets string
	}

	DatabaseCfg struct {
		Host string
		Port string
		User string
		Pass string
		Name string
		App  string
	}

	Logger struct {
		Info  LoggerSetting
		Error LoggerSetting
	}

	LoggerSetting struct {
		PathLog    string
		MaxSize    string //Megabytes
		MaxBackups string
		MaxAge     string //days
	}

	
)

var (
	App = appCfg{}      // Configuration Application
	PG  = DatabaseCfg{} // Configuration Database Postgresql
	LOG = Logger{}      // Configuration ElasticSeach
	Sys = SystemCfg{}
)

func GetEnv(Key, FallBack string) string {
	if Value, Ok := os.LookupEnv(Key); Ok {
		return Value
	}

	return FallBack
}

func (c *cfg) SetApp() appCfg {
	App.Host = GetEnv("APP_HOST", "localhost")
	App.Port = GetEnv("APP_PORT", "7951")
	App.Prefix = GetEnv("APP_PREFIX", "/playtorium")
	debug, _ := strconv.ParseBool(GetEnv("APP_DEBUG", "false"))
	App.Debug = debug
	return App
}

