package chrono

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// Date is used to save and output YYYY-MM-DD format date string.
type Date struct {
	time.Time
}

func (d Date) String() string {
	return d.In(time.UTC).Format(SQLDate)
}

// MarshalJSON converts a Time struct to ISO8601 string.
func (d Date) MarshalJSON() ([]byte, error) {

	if y := d.Year(); y < 0 || y >= 10000 {
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}
	if d.IsZero() {
		return []byte("null"), nil
	}

	b := make([]byte, 0, len(SQLDate)+2)
	b = append(b, '"')
	b = d.AppendFormat(b, SQLDate)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON converts ISO8601 data time into a Time struct.
// Empty string and null will be turned into time.Time zero value.
func (d *Date) UnmarshalJSON(data []byte) (err error) {

	if string(data) == "null" {
		d.Time = time.Time{}
		return
	}

	d.Time, err = time.Parse(`"`+SQLDate+`"`, string(data))

	return
}

// Scan implements the Scanner interface.
// SQL NULL will be turned into time zero value.
func (d *Date) Scan(value interface{}) (err error) {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		d.Time = v
		return
	case []byte:
		d.Time, err = time.Parse(SQLDate, string(v))
		return
	case string:
		d.Time, err = time.Parse(SQLDate, v)
		return
	}

	return fmt.Errorf("can't convert %T to time.Time", value)
}

// Value implements the driver Valuer interface.
// Zero value is turned into SQL NULL.
func (d Date) Value() (driver.Value, error) {
	if d.IsZero() {
		return nil, nil
	}

	return d.In(time.UTC).Format(SQLDate), nil
}

// DateNow creates current time.
func DateNow() Date {
	return Date{
		time.Now().Truncate(24 * time.Hour),
	}
}

func DateUTCNow() Date {
	return Date{
		time.Now().UTC().Truncate(24 * time.Hour),
	}
}

// DateZero creates the zero value of Time.
func DateZero() Date {
	return Date{time.Time{}}
}

// DateFrom creates a new Time wrapping time.Time.
func DateFrom(t time.Time) Date {
	return Date{t.Truncate(24 * time.Hour)}
}

func DateUTCFrom(t time.Time) Date {
	return Date{
		t.UTC().Truncate(24 * time.Hour),
	}
}
