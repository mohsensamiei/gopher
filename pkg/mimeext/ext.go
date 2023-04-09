package mimeext

import (
	"github.com/pinosell/gopher/pkg/slices"
	"mime"
)

var (
	routineExtensions = []string{
		".jpg", ".png", ".gif", ".svg", ".bmp",
		".mp4", "mov",
		".txt",
		".json", ".yaml", ".html",
	}
)

func ExtensionByType(typ string) (string, error) {
	list, err := mime.ExtensionsByType(typ)
	if err != nil {
		return "", err
	}
	for _, ext := range list {
		if slices.Contains(ext, routineExtensions...) {
			return ext, nil
		}
	}
	return list[0], nil
}
