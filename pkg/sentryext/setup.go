package sentryext

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	sentrylogrus "github.com/getsentry/sentry-go/logrus"
	"github.com/mohsensamiei/gopher/v2/pkg/slices"
	log "github.com/sirupsen/logrus"
	"time"
)

func Setup(configs Configs, service, version string) error {
	if !slices.Contains(configs.EnvName, configs.SentryEnvNames...) {
		return nil
	}
	client, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:              configs.SentryDSN,
		Debug:            configs.SentryDebug,
		AttachStacktrace: configs.SentryStack,
		EnableTracing:    configs.SentryTrace,
		SampleRate:       configs.SentryRate,
		TracesSampleRate: configs.SentryRate,
		Release:          fmt.Sprint(service, "@", version),
		Environment:      string(configs.EnvName),
	})
	if err != nil {
		return err
	}
	sentry.CurrentHub().BindClient(client)

	hook := sentrylogrus.NewFromClient([]log.Level{
		log.ErrorLevel,
		log.PanicLevel,
		log.FatalLevel,
	}, client)
	log.AddHook(hook)
	log.RegisterExitHandler(func() {
		hook.Flush(5 * time.Second)
	})
	return nil
}
