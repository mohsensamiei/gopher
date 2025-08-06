package timeext

import "time"

func FormatDuration(duration time.Duration, format string) string {
	return time.Unix(0, 0).UTC().Add(duration).Format(format)
}
