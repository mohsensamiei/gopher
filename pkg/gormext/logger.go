package gormext

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type logger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
}

func NewLogger() *logger {
	return &logger{
		SkipErrRecordNotFound: true,
	}
}

func (l *logger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *logger) Info(ctx context.Context, s string, args ...any) {
	log.WithContext(ctx).Infof(s, args...)
}

func (l *logger) Warn(ctx context.Context, s string, args ...any) {
	log.WithContext(ctx).Warnf(s, args...)
}

func (l *logger) Error(ctx context.Context, s string, args ...any) {
	log.WithContext(ctx).Errorf(s, args...)
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := log.Fields{
		"elapsed": fmt.Sprintf("%s", elapsed),
	}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[log.ErrorKey] = err
		log.WithContext(ctx).WithFields(fields).Errorf("%s", sql)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		log.WithContext(ctx).WithFields(fields).Warnf("%s", sql)
		return
	}

	log.WithContext(ctx).WithFields(fields).Debugf("%s", sql)
}
