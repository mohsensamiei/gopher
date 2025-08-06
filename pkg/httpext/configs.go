package httpext

import "github.com/mohsensamiei/gopher/v3/pkg/netext"

type Configs struct {
	HttpPort             netext.Port `env:"HTTP_PORT" envDefault:"8080"`
	FileRequestMaxSizeMB int64       `env:"FILE_REQUEST_MAX_SIZE_MB" envDefault:"5"`
	FileMaxSizeMB        int64       `env:"FILE_MAX_SIZE_MB" envDefault:"5"`
	FileAcceptMIMEList   []string    `env:"FILE_ACCEPT_MIME_LIST" envDefault:"application/pdf,image/jpeg,image/png"`
}

func (c Configs) FileRequestMaxSize() int64 {
	return c.FileRequestMaxSizeMB * 1024 * 1024
}

func (c Configs) FileMaxSize() int64 {
	return c.FileMaxSizeMB * 1024 * 1024
}
