package envext

type Mode int32

const (
	LOCAL Mode = iota
	DEVELOPMENT
	STAGING
	PRODUCTION
)
