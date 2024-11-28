package telegram

import "time"

type Date int64

func (d Date) Time() time.Time {
	return time.Unix(int64(d), 0)
}
