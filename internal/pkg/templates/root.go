package templates

const (
	RootMakefile = `
GOPATH=${HOME}/go

%:
	@true
`
	ConfigEnv = `
LOG_LEVEL=TRACE
`
)
