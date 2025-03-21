package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/gorm"
	"strconv"

)

var (
	DBConn *gorm.DB
	ZAPLog *zap.Logger
)

type AppConfig struct {
	App      *fiber.App
	Cfg      appCfg
	Logger   Logger
	Database DatabaseCfg
	System   SystemCfg
}

func NewApp() *AppConfig {
	return &AppConfig{}
}

func (c *AppConfig) SetApp() {
	c.App = fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	c.SetEnv()
	c.InitializeLogger()

	cfg := cors.Config{
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}

	c.App.Use(recover.New())
	c.App.Use(helmet.New())
	c.App.Use(cors.New(cfg))
	c.App.Use(logger.New(logger.Config{
		Format:     "${blue}${time} ${yellow}${status} - ${red}${latency} ${cyan}${method} ${path} ${green} ${ip} ${ua} ${reset}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Bangkok",
		Output:     os.Stdout,
	}))
}

func (c *AppConfig) SetEnv() {
	p := cfg{}
	c.Cfg = p.SetApp()
}

func (c *AppConfig) RunApp() {
	if err := c.App.Listen(fmt.Sprintf(":%v", c.Cfg.Port)); err != nil {
		logrus.Errorln("error app failed to start : ", err)
		return
	}
}


func (c *AppConfig) InitializeLogger() {
	stdout := zapcore.AddSync(os.Stdout)

	InfoMaxSize, _ := strconv.Atoi(LOG.Info.MaxSize)
	InfoMaxBackup, _ := strconv.Atoi(LOG.Info.MaxBackups)
	InfoMaxAge, _ := strconv.Atoi(LOG.Info.MaxAge)
	ErrorMaxSize, _ := strconv.Atoi(LOG.Error.MaxSize)
	ErrorMaxBackup, _ := strconv.Atoi(LOG.Error.MaxBackups)
	ErrorMaxAge, _ := strconv.Atoi(LOG.Error.MaxAge)

	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   LOG.Info.PathLog,
		MaxSize:    InfoMaxSize, // megabytes
		MaxBackups: InfoMaxBackup,
		MaxAge:     InfoMaxAge, // days
	})
	
	fileError := zapcore.AddSync(&lumberjack.Logger{
		Filename:   LOG.Error.PathLog,
		MaxSize:    ErrorMaxSize,   // megabytes -> 10
		MaxBackups: ErrorMaxBackup, // -> 3
		MaxAge:     ErrorMaxAge,    // days -> 45
	})

	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	levelError := zap.NewAtomicLevelAt(zap.ErrorLevel)

	//Set Zaplogger for write file
	productionCfg := zap.NewProductionEncoderConfig()
	productionCfg.TimeKey = "timestamp"
	productionCfg.MessageKey = "subject"
	productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	//Set Zaplogger for write console
	developmentCfg := zap.NewDevelopmentEncoderConfig()
	developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)
	fileEncoder := zapcore.NewJSONEncoder(productionCfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, stdout, level),
		zapcore.NewCore(fileEncoder, file, level),
		zapcore.NewCore(fileEncoder, fileError, levelError),
	)
	ZAPLog = zap.New(core)
}
