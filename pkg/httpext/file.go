package httpext

import (
	goerror "errors"
	"github.com/mohsensamiei/gopher/v3/pkg/errors"
	"github.com/mohsensamiei/gopher/v3/pkg/slices"
	"google.golang.org/grpc/codes"
	"io"
	"mime/multipart"
	"net/http"
)

type File struct {
	Name    string
	Size    int64
	MIME    string
	Content []byte
}

func ExtractFiles(configs Configs, req *http.Request) ([]*File, error) {
	if err := parseMultipartForm(configs, req); err != nil {
		return nil, err
	}
	var files []*File
	for name, _ := range req.MultipartForm.File {
		file, err := extractFile(configs, req, name)
		if err != nil {
			return nil, errors.Wrap(err, codes.InvalidArgument)
		}
		files = append(files, file)
	}
	return files, nil
}

func ExtractFile(configs Configs, req *http.Request, name string) (*File, error) {
	if err := parseMultipartForm(configs, req); err != nil {
		return nil, err
	}
	return extractFile(configs, req, name)
}

func parseMultipartForm(configs Configs, req *http.Request) error {
	if err := req.ParseMultipartForm(configs.FileRequestMaxSize()); err != nil {
		if goerror.Is(err, http.ErrNotMultipart) {
			return errors.WrapWithSlug(err, codes.Aborted, "invalid_file_size")
		}
		if goerror.Is(err, multipart.ErrMessageTooLarge) {
			return errors.WrapWithSlug(err, codes.ResourceExhausted, "invalid_file_size")
		}
		return errors.Wrap(err, codes.InvalidArgument)
	}
	return nil
}

func extractFile(configs Configs, req *http.Request, name string) (*File, error) {
	src, srcHeader, err := req.FormFile(name)
	if err != nil {
		if goerror.Is(err, http.ErrMissingFile) {
			return nil, errors.Wrap(err, codes.NotFound)
		}
		return nil, errors.New(codes.InvalidArgument).
			WithDetails(err.Error())
	}
	defer func() {
		_ = src.Close()
	}()
	file := &File{
		Name: srcHeader.Filename,
		Size: srcHeader.Size,
		MIME: srcHeader.Header.Get(ContentTypeHeader),
	}
	if !slices.Contains(file.MIME, configs.FileAcceptMIMEList...) {
		return nil, errors.NewWithSlug(codes.InvalidArgument, "invalid_file_type")
	}
	if file.Size > configs.FileMaxSize() {
		return nil, errors.NewWithSlug(codes.InvalidArgument, "invalid_file_size")
	}
	file.Content, err = io.ReadAll(src)
	if err != nil {
		return nil, errors.New(codes.InvalidArgument).
			WithDetails(err.Error())
	}
	return file, nil
}
