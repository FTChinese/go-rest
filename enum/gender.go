package enum

import (
	"database/sql/driver"
	"encoding/json"
)

// Gender values.
const (
	GenderZero   Gender = 0
	GenderFemale Gender = 1
	GenderMale   Gender = 2
)

var gendersRaw = [...]string{
	"",
	"F",
	"M",
}

// Gender is an enum for gender.
type Gender int

func (g Gender) String() string {
	if g < GenderFemale || g > GenderMale {
		return ""
	}

	return gendersRaw[g]
}

// UnmarshalJSON implements the Unmarshaler interface.
func (g *Gender) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	*g = NewGender(s)

	return nil
}

// MarshalJSON implmenets the Marshaler interface
func (g Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

// Scan implements sql.Scanner interface to retrieve enum value from SQL.
func (g *Gender) Scan(src interface{}) error {
	if src == nil {
		*g = GenderZero
		return nil
	}

	switch s := src.(type) {
	case []byte:
		*g = NewGender(string(s))
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

// NewGender converts a string into a Gender type.
func NewGender(gender string) Gender {
	switch gender {
	case "F":
		return GenderFemale
	case "M":
		return GenderMale
	default:
		return GenderZero
	}
}