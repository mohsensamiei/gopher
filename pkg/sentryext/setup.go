package sentryext

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

func Setup(configs Configs, service, version string) error {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              configs.SentryDSN,
		AttachStacktrace: true,
		EnableTracing:    true,
		Debug:            true,
		TracesSampleRate: 1.0,
		SampleRate:       1.0,
		Release:          fmt.Sprint(service, "@", version),
		Environment:      string(configs.EnvName),
	}); err != nil {
		return err
	}
	log.AddHook(&Hook{LogLevels: []log.Level{
		log.WarnLevel,
		log.ErrorLevel,
		log.PanicLevel,
		log.FatalLevel,
	}})
	return nil
}
