package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Gender is an enum for gender.
type Gender int

// Gender values.
const (
	InvalidGender Gender = iota
	GenderFemale
	GenderMale
)

var genderNames = [...]string{
	"",
	"F",
	"M",
}

var genderMap = map[Gender]string{
	1: genderNames[1],
	2: genderNames[2],
}

var genderValue = map[string]Gender{
	genderNames[1]: 1,
	genderNames[2]: 2,
}

// ParseGender parsed a string into Gender type.
func ParseGender(name string) (Gender, error) {
	if x, ok := genderValue[name]; ok {
		return x, nil
	}

	return InvalidGender, fmt.Errorf("%s is not a valid Gender", name)
}

func (g Gender) String() string {
	if s, ok := genderMap[g]; ok {
		return s
	}

	return ""
}

// UnmarshalJSON implements the Unmarshaler interface.
func (g *Gender) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tmp, err := ParseGender(s)

	if err != nil {
		return err
	}

	*g = tmp

	return nil
}

// MarshalJSON implmenets the Marshaler interface
func (g Gender) MarshalJSON() ([]byte, error) {
	s := g.String()
	if s == "" {
		return []byte("null"), nil
	}

	return []byte(`"` + s + `"`), nil
}

// Scan implements sql.Scanner interface to retrieve enum value from SQL.
func (g *Gender) Scan(src interface{}) error {
	if src == nil {
		*g = InvalidGender
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp, err := ParseGender(string(s))
		if err != nil {
			return err
		}
		*g = tmp
		return nil
	default:
		return ErrIncompatible
	}
}

// Value implements driver.Valuer interface to save value into SQL.
func (g Gender) Value() (driver.Value, error) {
	s := g.String()
	if s == "" {
		return nil, nil
	}

	return s, nil
}
