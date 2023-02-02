package templateext

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"text/template"
)

func toInt64(n any) int64 {
	switch v := n.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	default:
		panic(fmt.Sprintf("invalid input number '%v'", n))
	}
}

var (
	englishPrinter = message.NewPrinter(language.English)
	div            = template.FuncMap{"div": func(num, by any) int64 {
		return toInt64(num) / toInt64(by)
	}}
	cur = template.FuncMap{"cur": func(num any) string {
		return englishPrinter.Sprintf("%d", toInt64(num))
	}}
)

func Format(tmpl string, model any) (string, error) {
	tmp, err := template.New(uuid.New().String()).Funcs(div).Funcs(cur).Parse(tmpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err = tmp.Execute(&buf, model); err != nil {
		return "", err
	}
	return buf.String(), nil
}
