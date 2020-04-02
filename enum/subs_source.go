package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type SubsSource int

const (
	SubsSourceNull SubsSource = iota
	SubsSourceRetail
	SubsSourceB2B
)

var subsSourceNames = [...]string{
	"",
	"retail",
	"b2b",
}

// String representation of OrderKind
var subsSourceMap = map[SubsSource]string{
	1: subsSourceNames[1],
	2: subsSourceNames[2],
}

var subsSourceValue = map[string]SubsSource{
	subsSourceNames[1]: 1,
	subsSourceNames[2]: 2,
}

// ParseOrderKind creates OrderKind from a string.
func ParseSubsSource(name string) (SubsSource, error) {
	if x, ok := subsSourceValue[name]; ok {
		return x, nil
	}

	return SubsSourceNull, fmt.Errorf("%s is not valid OrderKind", name)
}

func (x SubsSource) String() string {
	if s, ok := subsSourceMap[x]; ok {
		return s
	}

	return ""
}

func (x *SubsSource) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tmp, _ := ParseSubsSource(s)

	*x = tmp

	return nil
}

func (x SubsSource) MarshalJSON() ([]byte, error) {
	s := x.String()

	if s == "" {
		return []byte("null"), nil
	}

	return []byte(`"` + s + `"`), nil
}

func (x *SubsSource) Scan(src interface{}) error {
	if src == nil {
		*x = SubsSourceNull
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp, _ := ParseSubsSource(string(s))
		*x = tmp
		return nil

	default:
		return errors.New("incompatible type to scan")
	}
}

func (x SubsSource) Value() (driver.Value, error) {
	s := x.String()
	if s == "" {
		return nil, nil
	}

	return s, nil
}
