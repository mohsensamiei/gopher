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
	GitIgnore = `
# IDE
.idea
.metals
.vscode

# OS files
.DS_Store
Thumbs.db

# Secret data
.env
*.env
*.pem

# Auto generated
/tmp
/logs
/build
/vendor
*.pb.go
docs.go
swagger.json
`
	DockerIgnore = `
# IDE
.idea
.metals
.vscode

# OS files
.DS_Store
Thumbs.db

# Secret data
.env
*.env
*.pem

# Auto generated
/tmp
/logs
/build
/vendor
*.pb.go
docs.go
swagger.json
`
	GitKeep       = ``
	MigrationUp   = ``
	MigrationDown = ``
	LanguageToml  = `
[unknown]
other = "..."

[internal]
other = "..."

[deadline_exceeded]
other = "..."

[already_exists]
other = "..."

[resource_exhausted]
other = "..."

[invalid_argument]
other = "..."

[not_found]
other = "..."

[failed_precondition]
other = "..."

[permission_denied]
other = "..."

[out_of_range]
other = "..."

[unimplemented]
other = "..."

[unavailable]
other = "..."

[unauthenticated]
other = "..."

[canceled]
other = "..."

[aborted]
other = "..."

[data_loss]
other = "..."
`
	DeployDockerfile = `
FROM ghcr.io/mohsensamiei/gopher/builder:latest as builder

WORKDIR /src
COPY go.mod go.sum ./
COPY vendor* vendor
RUN gopher dep
COPY . .
RUN gopher proto

# Add build lines here
# RUN GO111MODULE=on CGO_ENABLED=0 go build -buildvcs=false -a -installsuffix cgo \
#    -o ./build/OUTPUT ./cmd/MAIN

FROM ghcr.io/mohsensamiei/gopher/server:latest

WORKDIR /app
COPY --from=builder /src/build/ ./
COPY ./assets ./assets
`
	RootAir = `
[build]
  cmd = "sh -c 'go build -o tmp/output cmd/${SERVICE}/main.go'"
  bin = "tmp/output"
  delay = 1000
  exclude_dir = ["tmp", "vendor", "deploy", "scripts", "tests"]
  include_ext = ["go", "env", "tmpl", "toml", "sql", "pem"]
  exclude_regex = ["_test.go"]

[log]
  time = true

[misc]
  clean_on_exit = true
`
	CmdMain = `
package main

import (
	"github.com/mohsensamiei/gopher/v3/pkg/envext"
	"github.com/mohsensamiei/gopher/v3/pkg/health"
	"github.com/mohsensamiei/gopher/v3/pkg/i18next"
	"github.com/mohsensamiei/gopher/v3/pkg/logext"
	"github.com/mohsensamiei/gopher/v3/pkg/service"
	"github.com/mohsensamiei/gopher/v3/pkg/closer"
	log "github.com/sirupsen/logrus"
)

type Configs struct {
	Health   health.Configs
	Log      logext.Configs
	I18N     i18next.Configs
}

var (
	configs Configs
)

func init() {
	if err := envext.Parse(&configs); err != nil {
		log.WithError(err).Panic("can not parse env configs")
	}
	logext.Setup(configs.Log)
	if err := i18next.Setup(configs.I18N, "assets/locales"); err != nil {
		log.WithError(err).Panic("can not setup i18n package")
	}
}

func main() {
	defer closer.Defer()
	health.Serve(configs.Health)
	service.Start()
}
`
	ApiEnumImport = `
	"database/sql/driver"
	"encoding/json"
	"github.com/mohsensamiei/gopher/v3/pkg/mapext"`
	ApiEnum = `// region enum {{ .Enum }} methods
func ({{ .Enum }}) Values() []string {
	return mapext.Values({{ .Enum }}_name)
}
func ({{ .Enum }}) InRange(v interface{}) bool {
	_, ok := {{ .Enum }}_value[v.({{ .Enum }}).String()]
	return ok
}
func (x *{{ .Enum }}) Scan(value interface{}) error {
	*x = {{ .Enum }}({{ .Enum }}_value[value.(string)])
	return nil
}
func (x {{ .Enum }}) Value() (driver.Value, error) {
	return x.String(), nil
}
func (x *{{ .Enum }}) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*x = {{ .Enum }}({{ .Enum }}_value[str])
	return nil
}
func (x {{ .Enum }}) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
// endregion`
)
