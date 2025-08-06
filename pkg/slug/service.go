package slug

import "github.com/mohsensamiei/gopher/v3/pkg/stringsext"

type Service struct {
	Configs
}

func (s Service) Generate() string {
	return stringsext.Generate(s.SlugAlphabet, s.SlugLength)
}
