package s3ext

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/mohsensamiei/gopher/pkg/errors"
	"google.golang.org/grpc/codes"
	"net/http"
	"net/url"
	"strings"
)

func Dial(configs Configs) (*Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String(configs.S3EndpointURL.String()),
		Credentials: credentials.NewStaticCredentials(
			configs.S3AccessKey,
			configs.S3SecretKey,
			"",
		),
		HTTPClient: &http.Client{
			Timeout: configs.S3Timeout,
		},
		MaxRetries:       &configs.S3MaxRetries,
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, err
	}
	return &Client{
		Configs:    configs,
		session:    sess,
		s3:         s3.New(sess),
		uploader:   s3manager.NewUploader(sess),
		downloader: s3manager.NewDownloader(sess),
	}, nil
}

type Client struct {
	Configs
	s3         *s3.S3
	session    *session.Session
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
}

func (c *Client) Upload(filename string, data []byte, public bool) (*url.URL, error) {
	return c.UploadWithContext(context.Background(), filename, data, public)
}

func (c *Client) UploadWithContext(ctx context.Context, filepath string, data []byte, public bool) (*url.URL, error) {
	req := s3manager.UploadInput{
		Bucket: aws.String(c.S3BucketName),
		Key:    aws.String(filepath),
		Body:   bytes.NewReader(data),
		ACL:    aws.String("private"),
	}
	if public {
		req.ACL = aws.String("public-read")
	}
	res, err := c.uploader.UploadWithContext(ctx, &req)
	if err != nil {
		return nil, err
	}

	var location *url.URL
	location, err = url.Parse(res.Location)
	if err != nil {
		return nil, err
	}
	return location, nil
}

func (c *Client) DownloadByURLWithContext(ctx context.Context, uri url.URL) ([]byte, error) {
	if uri.Host != c.S3EndpointURL.Host {
		return nil, errors.New(codes.InvalidArgument).
			WithDetailF("invalid host '%v'", uri.Host)
	}
	dump := strings.Split(uri.Path, "/")
	if len(dump) < 2 {
		return nil, errors.New(codes.InvalidArgument).
			WithDetailF("invalid file path '%v'", uri.Path)
	}
	return c.DownloadWithContext(ctx, strings.Join(dump[2:], "/"))
}

func (c *Client) DownloadWithContext(ctx context.Context, filepath string) ([]byte, error) {
	buffer := aws.NewWriteAtBuffer(make([]byte, c.S3BufferSize*1024))
	// if the downloaded object is larger than our buffer, it will grow by this factor
	buffer.GrowthCoeff = c.S3BufferGrowth

	if _, err := c.downloader.DownloadWithContext(ctx, buffer, &s3.GetObjectInput{
		Bucket: aws.String(c.S3BucketName),
		Key:    aws.String(filepath),
	}); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (c *Client) DeleteWithContext(ctx context.Context, filepath string) error {
	if _, err := c.s3.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.S3BucketName),
		Key:    aws.String(filepath),
	}); err != nil {
		return err
	}
	return nil
}
