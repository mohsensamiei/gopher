package jwtext

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Configs struct {
	JwtPublicKey     string        `env:"JWT_PUBLIC_KEY,required,file"`
	JwtPrivateKey    string        `env:"JWT_PRIVATE_KEY,file"`
	JwtTokenDuration time.Duration `env:"JWT_TOKEN_DURATION" envDefault:"24h"`
}

func Setup(configs Configs) (*JWT, error) {
	var (
		err      error
		instance = &JWT{
			duration: configs.JwtTokenDuration,
		}
	)
	instance.verifyKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(configs.JwtPublicKey))
	if err != nil {
		return nil, err
	}
	if configs.JwtPrivateKey != "" {
		instance.signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(configs.JwtPrivateKey))
		if err != nil {
			return nil, err
		}
	}
	return instance, nil
}
