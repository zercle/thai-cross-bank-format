package utils

import (
	"errors"
	"time"
)

const (
	DateFmt      = "2006-01-02"
	TimeFmt      = "15:04:00"
	DateTimeFmt  = "2006-01-02 15:04:05"
	RFC3339Milli = "2006-01-02T15:04:05.999Z07:00"
)

type TimeMilli time.Time

// MarshalJSON implements the json.Marshaler interface.
// The time is a quoted string in RFC 3339 format, with sub-second precision added if present.
func (t TimeMilli) MarshalJSON() (buff []byte, err error) {
	if y := time.Time(t).Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	buff = make([]byte, 0, len(RFC3339Milli)+2)
	buff = append(buff, '"')
	buff = time.Time(t).AppendFormat(buff, RFC3339Milli)
	buff = append(buff, '"')
	return
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TimeMilli) UnmarshalJSON(data []byte) (err error) {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tmp, err := time.ParseInLocation(`"`+RFC3339Milli+`"`, string(data), time.Local)
	*t = TimeMilli(tmp)
	return
}

// MarshalText implements the encoding.TextMarshaler interface.
func (t TimeMilli) MarshalText() ([]byte, error) {
	if y := time.Time(t).Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalText: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(RFC3339Milli))
	return time.Time(t).AppendFormat(b, RFC3339Milli), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (t *TimeMilli) UnmarshalText(data []byte) (err error) {
	// Fractional seconds are handled implicitly by Parse.
	tmp, err := time.ParseInLocation(RFC3339Milli, string(data), time.Local)
	*t = TimeMilli(tmp)
	return
}
