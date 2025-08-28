package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"moul.io/http2curl/v2"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

type method string

func (m method) String() string {
	return string(m)
}

const (
	getMe                  method = "getMe"
	getFile                method = "getFile"
	getUpdates             method = "getUpdates"
	sendMessage            method = "sendMessage"
	sendChatAction         method = "sendChatAction"
	sendDocument           method = "sendDocument"
	editMessageText        method = "editMessageText"
	deleteMessage          method = "deleteMessage"
	createChatInviteLink   method = "createChatInviteLink"
	approveChatJoinRequest method = "approveChatJoinRequest"
	declineChatJoinRequest method = "declineChatJoinRequest"
	setWebhook             method = "setWebhook"
	deleteWebhook          method = "deleteWebhook"
)

const (
	uri = "https://api.telegram.org/bot%v/%v"
)

func request[T any](token string, method method, request any, response *Response[T]) error {
	body, content := multipartBody(request)
	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf(uri, token, method),
		body,
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", content)

	command, _ := http2curl.GetCurlCommand(req)
	log.WithField("curl", command.String()).Trace("telegram api request")

	var res *http.Response
	res, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode != 200 {
		var (
			bin  []byte
			data = new(Error)
		)
		bin, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(bin, data); err != nil {
			return err
		}
		return fmt.Errorf("%v: %v", data.ErrorCode, data.Description)
	}

	var bin []byte
	bin, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(bin, response); err != nil {
		return err
	}
	return nil
}

func multipartBody(s any) (*bytes.Buffer, string) {
	zeroBuf := new(bytes.Buffer)
	zeroContent := "multipart/form-data"
	if s == nil {
		return zeroBuf, zeroContent
	}

	empty := true
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	structValue := reflect.ValueOf(s)
	structType := structValue.Type()
	for i := 0; i < structType.NumField(); i++ {
		fieldType := structType.Field(i)
		if !fieldType.IsExported() {
			continue
		}

		var (
			isOmitempty bool
			name        string
			qs          = strings.Split(fieldType.Tag.Get("json"), ",")
		)
		if l := len(qs); l == 1 {
			name = strings.TrimSpace(qs[0])
		} else if l == 2 {
			name = strings.TrimSpace(qs[0])
			isOmitempty = qs[1] == "omitempty"
		} else {
			name = fieldType.Name
		}

		fieldVal := structValue.Field(i)
		if isOmitempty && fieldVal.IsZero() {
			continue
		}

		empty = false
		switch v := fieldVal.Interface().(type) {
		case InputFileContent:
			part, _ := writer.CreateFormFile(name, filepath.Base(v.Name))
			_, _ = io.Copy(part, bytes.NewReader(v.Content))
		case InputFileID:
			_ = writer.WriteField(name, v.ID)
		case time.Duration:
			_ = writer.WriteField(name, fmt.Sprint(v.Seconds()))
		default:
			switch fieldType.Type.Kind() {
			case reflect.Interface,
				reflect.Struct,
				reflect.Array,
				reflect.Map:
				j, _ := json.Marshal(fieldVal.Interface())
				_ = writer.WriteField(name, fmt.Sprintf("%s", j))
			default:
				_ = writer.WriteField(name, fmt.Sprint(v))
			}
		}
	}
	_ = writer.Close()

	if empty {
		return new(bytes.Buffer), "multipart/form-data"
	}
	return body, writer.FormDataContentType()
}
