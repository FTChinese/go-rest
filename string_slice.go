package gorest

import (
	"database/sql/driver"
	"errors"
	"strings"
)

// StringSlice implements Scanner and Valuer interface for a slice of strings.
type StringSlice []string

// Scan retrieves a comma-separated string value from SQL to a Go string slice.
func (x *StringSlice) Scan(src interface{}) error {
	if src == nil {
		*x = []string{}
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp := strings.Split(string(s), ",")
		*x = tmp
		return nil

	default:
		return errors.New("incompatible type to scan")
	}
}

// Value turns a Go string slice to a comma-sparated string.
func (x StringSlice) Value() (driver.Value, error) {
	s := strings.Join(x, ",")
	if s == "" {
		return nil, nil
	}

	return s, nil
}
