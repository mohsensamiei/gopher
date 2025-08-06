package logext

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type Configs struct {
	LogLevel log.Level `env:"LOG_LEVEL" envDefault:"INFO"`
	LogFile  string    `env:"LOG_FILE"`
	LogSize  int       `env:"LOG_SIZE" envDefault:"100"` // megabytes
	LogAge   int       `env:"LOG_AGE" envDefault:"30"`   // days
}

func Setup(configs Configs) {
	log.SetLevel(configs.LogLevel)
	if configs.LogFile != "" {
		log.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
			Filename: configs.LogFile,
			MaxSize:  configs.LogSize,
			MaxAge:   configs.LogAge,
		}))
	}
}
