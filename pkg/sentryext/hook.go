package sentryext

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"github.com/mohsensamiei/gopher/v3/pkg/stringsext"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

var (
	level = map[log.Level]sentry.Level{
		log.DebugLevel: sentry.LevelDebug,
		log.InfoLevel:  sentry.LevelInfo,
		log.WarnLevel:  sentry.LevelWarning,
		log.ErrorLevel: sentry.LevelError,
		log.FatalLevel: sentry.LevelFatal,
		log.PanicLevel: sentry.LevelFatal,
		log.TraceLevel: sentry.LevelDebug,
	}
)

type Hook struct {
	LogLevels []log.Level
}

func (hook *Hook) Fire(entry *log.Entry) error {
	if strings.Contains(stringsext.Comparable(entry.Message), "context canceled") {
		return nil
	}
	defer sentry.Flush(5 * time.Second)
	event := &sentry.Event{
		Message: entry.Message,
		Level:   level[entry.Level],
		Exception: []sentry.Exception{
			{
				Type:  entry.Message,
				Value: strings.ToLower(entry.Level.String()),
			},
		},
		Extra: make(map[string]interface{}),
		Threads: []sentry.Thread{
			{
				Stacktrace: NewStacktrace(3),
			},
		},
	}
	for k, v := range entry.Data {
		event.Extra[k] = fmt.Sprintf("%+v", v)
	}
	if v, ok := entry.Data[log.ErrorKey]; ok {
		err := errors.Cast(v.(error))
		event.Exception[0].Value = err.Message()
		event.Message = strings.Join(err.Details(), "\n")
	}
	sentry.CaptureEvent(event)
	return nil
}

// Levels define on which log levels this hook would trigger
func (hook *Hook) Levels() []log.Level {
	return hook.LogLevels
}
