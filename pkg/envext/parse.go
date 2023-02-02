package envext

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

var (
	defaultTypeParsers = map[reflect.Type]env.ParserFunc{
		reflect.TypeOf(true): func(v string) (any, error) {
			return strconv.ParseBool(strings.TrimSpace(v))
		},
		reflect.TypeOf(""): func(v string) (any, error) {
			return strings.TrimSpace(v), nil
		},
		reflect.TypeOf(1): func(v string) (any, error) {
			i, err := strconv.ParseInt(strings.TrimSpace(v), 10, 32)
			return int(i), err
		},
		reflect.TypeOf(int16(1)): func(v string) (any, error) {
			i, err := strconv.ParseInt(strings.TrimSpace(v), 10, 16)
			return int16(i), err
		},
		reflect.TypeOf(int32(1)): func(v string) (any, error) {
			i, err := strconv.ParseInt(strings.TrimSpace(v), 10, 32)
			return int32(i), err
		},
		reflect.TypeOf(int64(1)): func(v string) (any, error) {
			return strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		},
		reflect.TypeOf(int8(1)): func(v string) (any, error) {
			i, err := strconv.ParseInt(strings.TrimSpace(v), 10, 8)
			return int8(i), err
		},
		reflect.TypeOf(uint(1)): func(v string) (any, error) {
			i, err := strconv.ParseUint(strings.TrimSpace(v), 10, 32)
			return uint(i), err
		},
		reflect.TypeOf(uint16(1)): func(v string) (any, error) {
			i, err := strconv.ParseUint(strings.TrimSpace(v), 10, 16)
			return uint16(i), err
		},
		reflect.TypeOf(uint32(1)): func(v string) (any, error) {
			i, err := strconv.ParseUint(strings.TrimSpace(v), 10, 32)
			return uint32(i), err
		},
		reflect.TypeOf(uint64(1)): func(v string) (any, error) {
			i, err := strconv.ParseUint(strings.TrimSpace(v), 10, 64)
			return i, err
		},
		reflect.TypeOf(uint8(1)): func(v string) (any, error) {
			i, err := strconv.ParseUint(strings.TrimSpace(v), 10, 8)
			return uint8(i), err
		},
		reflect.TypeOf(float64(1)): func(v string) (any, error) {
			return strconv.ParseFloat(strings.TrimSpace(v), 64)
		},
		reflect.TypeOf(float32(1)): func(v string) (any, error) {
			f, err := strconv.ParseFloat(strings.TrimSpace(v), 32)
			return float32(f), err
		},
		reflect.TypeOf(url.URL{}): func(v string) (any, error) {
			u, err := url.Parse(strings.TrimSpace(v))
			if err != nil {
				return nil, fmt.Errorf("unable to parse URL: %v", err)
			}
			return *u, nil
		},
		reflect.TypeOf(time.Nanosecond): func(v string) (any, error) {
			s, err := time.ParseDuration(strings.TrimSpace(v))
			if err != nil {
				return nil, fmt.Errorf("unable to parse duration: %v", err)
			}
			return s, err
		},
		reflect.TypeOf(make(map[string]string)): func(v string) (any, error) {
			dic := make(map[string]string)
			for _, i := range strings.Split(v, ",") {
				dump := strings.Split(i, ":")
				if len(dump) != 2 {
					return nil, fmt.Errorf("unable to parse dictionary")
				}
				dic[dump[0]] = dump[1]
			}
			return dic, nil
		},
	}
)

func Parse(configs any) error {
	if err := env.ParseWithFuncs(configs, defaultTypeParsers); err != nil {
		return err
	}
	return nil
}
