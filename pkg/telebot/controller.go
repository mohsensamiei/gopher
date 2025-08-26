package telebot

import (
	"github.com/gorilla/mux"
	"github.com/mohsensamiei/gopher/v3/pkg/authenticate"
	"github.com/mohsensamiei/gopher/v3/pkg/httpext"
	"github.com/mohsensamiei/gopher/v3/pkg/muxext"
	"github.com/mohsensamiei/gopher/v3/pkg/stringsext"
	"github.com/mohsensamiei/gopher/v3/pkg/telegram"
	"net/http"
)

func (c *Client) NewController() muxext.ControllerRegister {
	return &Controller{
		Client: c,
	}
}

const (
	secretTokenHeader = "X-Telegram-Bot-Api-Secret-Token"
)

type Controller struct {
	*Client
}

func (c Controller) RegisterController(router *mux.Router) {
	muxext.HandleFunc(router, "/v1/hooks/telegram", c.Update).Methods(http.MethodPost)
}

func (c Controller) Update(res http.ResponseWriter, req *http.Request) {
	if !stringsext.IsNilOrEmpty(c.TelegramSecretToken) && httpext.Header(req, secretTokenHeader) != c.TelegramSecretToken {
		httpext.SendError(res, req, authenticate.ErrUnauthenticated)
		return
	}
	var update telegram.Update
	if err := httpext.BindRequestModel(req, &update); err != nil {
		httpext.SendError(res, req, err)
		return
	}
	c.channelUpdateId = update.UpdateID + 1
	c.channel <- update
	httpext.SendCode(res, req, http.StatusOK)
}
