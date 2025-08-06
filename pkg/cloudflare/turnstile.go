package cloudflare

import (
	"bytes"
	"encoding/json"
	"github.com/go-redis/redis_rate/v10"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mohsensamiei/gopher/v3/pkg/di"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"github.com/mohsensamiei/gopher/v3/pkg/httpext"
	"github.com/mohsensamiei/gopher/v3/pkg/mimeext"
	"github.com/mohsensamiei/gopher/v3/pkg/redisext"
	"google.golang.org/grpc/codes"
	"net/http"
	"strings"
	"time"
)

const (
	turnstileApi = "https://challenges.cloudflare.com/turnstile/v0/siteverify"
)

type TurnstileRequest struct {
	Secret         string `json:"secret"`
	Response       string `json:"response"`
	RemoteIP       string `json:"remoteip"`
	IdempotencyKey string `json:"idempotency_key"`
}

type TurnstileResponse struct {
	Success     bool      `json:"success"`
	ChallengeTs time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
	Action      string    `json:"action"`
	CData       string    `json:"cdata"`
}

type Turnstile struct {
	TurnstileSecretKey string `env:"CF_TURNSTILE_SECRET_KEY,required"`
}

func (t *Turnstile) Verify(response, remoteIP string) error {
	payload, _ := json.Marshal(TurnstileRequest{
		Secret:         t.TurnstileSecretKey,
		Response:       response,
		RemoteIP:       remoteIP,
		IdempotencyKey: uuid.New().String(),
	})
	res, err := http.DefaultClient.Post(turnstileApi, mimeext.Json, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()
	result := new(TurnstileResponse)
	if _, err = httpext.BindResponseModel(res, result); err != nil {
		return err
	}
	if !result.Success {
		return errors.New(codes.PermissionDenied).WithDetailF(strings.Join(result.ErrorCodes, ","))
	}
	return nil
}

func TurnstileMiddleware(hourLimit int) mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			remoteIP := httpext.Header(req, "CF-Connecting-IP")
			response := httpext.Header(req, "CF-Turnstile-Response")
			turnstile := di.Provide[*Turnstile](req.Context())
			if hourLimit > 0 {
				if r, err := di.Provide[*redisext.Client](req.Context()).Limit(req.Context(), "turnstile:ip", remoteIP,
					redis_rate.PerHour(hourLimit),
				); err != nil {
					httpext.SendError(res, req, err)
					return
				} else if r.Allowed <= 0 {
					httpext.SendError(res, req, errors.New(codes.ResourceExhausted))
					return
				}
				if err := turnstile.Verify(response, remoteIP); err != nil {
					httpext.SendError(res, req, err)
					return
				}
			}
			handler.ServeHTTP(res, req)
		})
	}
}
