package logext

import (
	"github.com/mohsensamiei/gopher/v3/pkg/mapext"
	log "github.com/sirupsen/logrus"
)

func newFormatter(defaults log.Fields) log.Formatter {
	return &formatter{
		JSONFormatter: new(log.JSONFormatter),
		DefaultFields: defaults,
	}
}

type formatter struct {
	*log.JSONFormatter
	DefaultFields log.Fields
}

func (f *formatter) Format(entry *log.Entry) ([]byte, error) {
	entry.Data = mapext.Merge(entry.Data, f.DefaultFields)
	return f.JSONFormatter.Format(entry)
}
