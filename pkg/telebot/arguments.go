package telebot

import (
	"errors"
	"strings"
)

func NewArguments() Arguments {
	return make(Arguments)
}

type Arguments map[string][]string

func (a *Arguments) Parse(s string) error {
	*a = make(Arguments)
	if s == "" {
		return nil
	}
	parts := strings.Split(s, "__")
	for _, p := range parts {
		kv := strings.SplitN(p, "_", 2)
		if len(kv) != 2 {
			return errors.New("invalid format")
		}
		key := kv[0]
		values := strings.Split(kv[1], "-")
		(*a)[key] = values
	}
	return nil
}

func (a Arguments) String() string {
	parts := make([]string, 0, len(a))
	for k, vals := range a {
		parts = append(parts, k+"_"+strings.Join(vals, "-"))
	}
	return strings.Join(parts, "__")
}

func (a Arguments) Exists(key string) bool {
	_, ok := a[key]
	return ok
}

func (a Arguments) Get(key string) ([]string, bool) {
	vals, ok := a[key]
	return vals, ok
}

func (a Arguments) Set(key string, values ...string) {
	a[key] = values
}

func (a Arguments) Add(key string, values ...string) {
	a[key] = append(a[key], values...)
}
