package chrono

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// Time is used to save and output ISO8601 date time with time zone set to UTC.
type Time struct {
	time.Time
}

// StringEN produces the string representation in English with locale set to UTC.
func (t Time) StringEN() string {
	return t.In(time.UTC).Format(time.RFC1123Z)
}

// StringCN produces the string representation in Chinese format with locale set to Asia/Shanghai.
func (t Time) StringCN() string {
	return t.In(TZShanghai).Format(CST)
}

// MarshalJSON converts a Time struct to ISO8601 string.
func (t Time) MarshalJSON() ([]byte, error) {
	if y := t.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}
	if t.IsZero() {
		return []byte("null"), nil
	}


	b := make([]byte, 0, len(time.RFC3339)+2)
	b = append(b, '"')
	b = t.AppendFormat(b, time.RFC3339)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The time is expected to be a quoted string in RFC 3339 format.
// Empty string and null will be turned into time.Time zero value.
func (t *Time) UnmarshalJSON(data []byte) (err error) {

	if string(data) == "null" {
		t.Time = time.Time{}
		return
	}

	t.Time, err = time.Parse(`"`+time.RFC3339+`"`, string(data))

	return
}

// Scan implements the Scanner interface.
// SQL NULL will be turned into time zero value.
func (t *Time) Scan(value interface{}) (err error) {
	if value == nil {
		t.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return
	case []byte:
		t.Time, err = ParseDateTime(string(v), time.UTC)
		return
	case string:
		t.Time, err = ParseDateTime(v, time.UTC)
		return
	}

	return fmt.Errorf("can't convert %T to time.Time", value)
}

// Value implements the driver Valuer interface.
// Zero value is turned into SQL NULL.
func (t Time) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}

	return t.In(time.UTC).Format(SQLDateTime), nil
}

// TimeNow creates current time.
func TimeNow() Time {
	return Time{
		time.Now(),
	}
}

// TimeZero creates the zero value of Time.
func TimeZero() Time {
	return Time{time.Time{}}
}

// TimeFrom creates a new Time wrapping time.Time.
func TimeFrom(t time.Time) Time {
	return Time{t}
}