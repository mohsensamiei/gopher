package sentryext

import "github.com/mohsensamiei/gopher/v2/pkg/envext"

type Configs struct {
	EnvName   envext.Name `env:"ENV_NAME" envDefault:"LOCAL"`
	SentryDSN string      `env:"SENTRY_DSN,required"`
}
