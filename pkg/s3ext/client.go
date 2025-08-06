package s3ext

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"google.golang.org/grpc/codes"
	"net/http"
	"net/url"
	"strings"
)

func Dial(ctx context.Context, configs Configs) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-west-2"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				configs.S3AccessKey,
				configs.S3SecretKey,
				"",
			),
		),
		config.WithHTTPClient(&http.Client{
			Timeout: configs.S3Timeout,
		}),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				if service == s3.ServiceID {
					return aws.Endpoint{
						URL:           configs.S3EndpointURL.String(),
						SigningRegion: region,
					}, nil
				}
				return aws.Endpoint{}, &aws.EndpointNotFoundError{}
			}),
		),
	)
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.Retryer = retry.AddWithMaxAttempts(retry.NewStandard(), configs.S3MaxRetries)
	})
	return &Client{
		Configs:    configs,
		S3Client:   s3Client,
		Uploader:   manager.NewUploader(s3Client),
		Downloader: manager.NewDownloader(s3Client),
	}, nil
}

type Client struct {
	Configs
	S3Client   *s3.Client
	Uploader   *manager.Uploader
	Downloader *manager.Downloader
}

func (c *Client) Upload(filename string, data []byte, public bool) (*url.URL, error) {
	return c.UploadWithContext(context.Background(), filename, data, public)
}
func (c *Client) UploadWithContext(ctx context.Context, filepath string, data []byte, public bool) (*url.URL, error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(c.Configs.S3BucketName),
		Key:    aws.String(filepath),
		Body:   bytes.NewReader(data),
		ACL:    types.ObjectCannedACLPrivate,
	}
	if public {
		input.ACL = types.ObjectCannedACLPublicRead
	}

	var location *url.URL
	if res, err := c.Uploader.Upload(ctx, input); err != nil {
		return nil, err
	} else if location, err = url.Parse(res.Location); err != nil {
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
	buf := manager.NewWriteAtBuffer([]byte{})
	if _, err := c.Downloader.Download(ctx, buf, &s3.GetObjectInput{
		Bucket: aws.String(c.Configs.S3BucketName),
		Key:    aws.String(filepath),
	}); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *Client) DeleteWithContext(ctx context.Context, filepath string) error {
	if _, err := c.S3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(c.S3BucketName),
		Key:    aws.String(filepath),
	}); err != nil {
		return err
	}
	return nil
}
