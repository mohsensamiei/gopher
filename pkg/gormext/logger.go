package gormext

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	gormlogger "gorm.io/gorm/logger"
	"time"
)

type Logger struct{}

func NewLogger() *Logger {
	return new(Logger)
}

func (l *Logger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *Logger) Info(ctx context.Context, s string, args ...any) {
	log.WithContext(ctx).Infof(s, args...)
}

func (l *Logger) Warn(ctx context.Context, s string, args ...any) {
	log.WithContext(ctx).Warnf(s, args...)
}

func (l *Logger) Error(ctx context.Context, s string, args ...any) {
	log.WithContext(ctx).Errorf(s, args...)
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	fields := log.Fields{
		"elapsed": fmt.Sprintf("%s", time.Since(begin)),
	}
	if err != nil {
		fields[log.ErrorKey] = err
	}
	sql, _ := fc()
	log.WithContext(ctx).WithFields(fields).Tracef("%s", sql)
}
