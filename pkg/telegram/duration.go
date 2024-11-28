package telegram

import "time"

type Duration int

func (d Duration) Duration() time.Duration {
	return time.Duration(d) * time.Second
}
