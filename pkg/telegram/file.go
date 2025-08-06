package telegram

import (
	"bytes"
	"fmt"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"google.golang.org/grpc/codes"
	"io"
	"net/http"
)

// File
// This object represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
// The maximum file size to download is 20 MB
type File struct {
	FileID       string `json:"file_id"`        // Identifier for this file, which can be used to download or reuse the file
	FileUniqueID string `json:"file_unique_id"` // Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int    `json:"file_size"`      // Optional. File size in bytes. It can be bigger than 2^31 and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value.
	FilePath     string `json:"file_path"`      // Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

type GetFile struct {
	FileID string `json:"file_id"` // Required. File identifier to get information about
}

// GetFile
// Use this method to get basic information about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
func (c Connection) GetFile(req GetFile) (*File, io.Reader, error) {
	var res Response[File]
	if err := request(c.TelegramToken, getFile, req, &res); err != nil {
		return nil, nil, err
	}

	var data []byte
	{
		response, err := http.Get(fmt.Sprintf("https://api.telegram.org/file/bot%v/%v", c.TelegramToken, res.Result.FilePath))
		if err != nil {
			return nil, nil, err
		}
		defer func() {
			_ = response.Body.Close()
		}()
		if response.StatusCode != http.StatusOK {
			return nil, nil, errors.New(codes.Internal).
				WithDetailF("can not download file: %v", response.Status)
		}

		data, err = io.ReadAll(response.Body)
		if err != nil {
			return nil, nil, err
		}
	}
	return &res.Result, bytes.NewReader(data), nil
}
