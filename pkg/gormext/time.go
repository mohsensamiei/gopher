package gormext

import (
	"database/sql/driver"
	"time"
)

type NullTime struct {
	time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	nt.Valid = nt.parsePGTime(value.(string)) == nil
	return nil
}

func (nt *NullTime) parsePGTime(val string) error {
	var err error
	nt.Time, err = time.Parse("15:04:05", val)
	if err != nil {
		return err
	}
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid { //t.Time.Format("15:04:05")
		return nil, nil
	}
	return nt.Time.Format("15:04:05"), nil
}

func (nt *NullTime) Set(t *time.Time) NullTime {
	if t != nil {
		nt.Time = *t
		nt.Valid = true
	}
	return *nt
}
