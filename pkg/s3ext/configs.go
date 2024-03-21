package s3ext

import (
	"net/url"
	"time"
)

type Configs struct {
	S3EndpointURL  url.URL       `env:"S3_ENDPOINT_URL,required"`
	S3AccessKey    string        `env:"S3_ACCESS_KEY,required"`
	S3SecretKey    string        `env:"S3_SECRET_KEY,required"`
	S3BucketName   string        `env:"S3_BUCKET_NAME,required"`
	S3Timeout      time.Duration `env:"S3_TIMEOUT" envDefault:"5s"`
	S3MaxRetries   int           `env:"S3_MAX_RETRIES" envDefault:"3"`
	S3BufferSize   int           `env:"S3_BUFFER_SIZE_KB" envDefault:"500"`
	S3BufferGrowth float64       `env:"S3_BUFFER_GROWTH" envDefault:"1.8"`
}
