package utils

import (
	"time"
)

const (
	dateDMY        = `02.01.2006`
	dateYMD        = `2006.01.02`
	timeHMS        = `15:04:05`
	datetimeDMYHMS = `02.01.2006 15:04:05`
	datetimeYMDHMS = `15:04:05 02.01.2006`
)

func ParseDate(in string) (time.Time, error) {
	data, err := time.Parse(dateDMY, in)
	if err == nil {
		return data, nil
	}
	return time.Parse(dateYMD, in)
}

func ParseTime(in string) (time.Time, error) {
	return time.Parse(timeHMS, in)
}

func ParseDatetime(in string) (time.Time, error) {
	var data time.Time
	var err error

	data, err = time.Parse(datetimeDMYHMS, in)
	if err == nil {
		return data, nil
	}
	data, err = time.Parse(datetimeYMDHMS, in)
	if err == nil {
		return data, nil
	}
	return time.Parse(time.RFC3339, in)
}
