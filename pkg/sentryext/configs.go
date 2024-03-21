package sentryext

import "github.com/mohsensamiei/gopher/v2/pkg/envext"

type Configs struct {
	EnvName        envext.Name   `env:"ENV_NAME" envDefault:"LOCAL"`
	SentryEnvNames []envext.Name `env:"SENTRY_ENV_NAMES" envDefault:"DEVELOPMENT,STAGING,PRODUCTION"`
	SentryDSN      string        `env:"SENTRY_DSN,required"`
	SentryDebug    bool          `env:"SENTRY_DEBUG" envDefault:"false"`
	SentryStack    bool          `env:"SENTRY_STACK" envDefault:"true"`
	SentryTrace    bool          `env:"SENTRY_TRACE" envDefault:"true"`
	SentryRate     float64       `env:"SENTRY_RATE" envDefault:"1.0"`
}
