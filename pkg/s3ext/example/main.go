package main

import (
	"context"
	"github.com/mohsensamiei/gopher/v3/pkg/envext"
	"github.com/mohsensamiei/gopher/v3/pkg/s3ext"
	"net/url"
	"os"
)

func init() {
	_ = os.Setenv("S3_ENDPOINT_URL", "https://dev-itbfx.fra1.digitaloceanspaces.com")
	_ = os.Setenv("S3_ACCESS_KEY", "DO801JXV8Y3MTN4WNYCY")
	_ = os.Setenv("S3_SECRET_KEY", "Df8xl5i1ZSW2lCe6C/r0NaBNTA3/oFqCPt/lZqpFhbg")
	_ = os.Setenv("S3_BUCKET_NAME", "dev-itbfx")
}

func main() {
	var configs s3ext.Configs
	if err := envext.Parse(&configs); err != nil {
		panic(err)
	}

	client, err := s3ext.Dial(context.Background(), configs)
	if err != nil {
		panic(err)
	}

	var uri *url.URL
	uri, err = client.Upload("docs/test.txt", []byte("hello world"), true)
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
