package main

import (
	"context"
	"github.com/mohsensamiei/gopher/v2/pkg/envext"
	"github.com/mohsensamiei/gopher/v2/pkg/s3ext"
	"net/url"
	"os"
)

func init() {
	_ = os.Setenv("S3_ENDPOINT_URL", "")
	_ = os.Setenv("S3_ACCESS_KEY", "")
	_ = os.Setenv("S3_SECRET_KEY", "")
}

func main() {
	var configs s3ext.Configs
	if err := envext.Parse(&configs); err != nil {
		panic(err)
	}

	client, err := s3ext.Dial(configs)
	if err != nil {
		panic(err)
	}

	var uri *url.URL
	uri, err = client.Upload("/docs/test.txt", []byte("hello world"), true)
	if err != nil {
		panic(err)
	}
	println(uri.String())

	var data []byte
	data, err = client.DownloadByURLWithContext(context.Background(), *uri)
	if err != nil {
		panic(err)
	}
	println(string(data))
}
