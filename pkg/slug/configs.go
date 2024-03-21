package slug

type Configs struct {
	SlugAlphabet string `env:"SLUG_ALPHABET" envDefault:"abcdefghijklmnopqrstuvwxyz0123456789"`
	SlugLength   int    `env:"SLUG_LENGTH" envDefault:"8"`
}

func NewService(configs Configs) *Service {
	return &Service{
		Configs: configs,
	}
}
