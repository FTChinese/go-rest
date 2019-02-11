package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// Cycle is an enum for billing cycles.
type Cycle int

// Supported billing cycles
const (
	InvalidCycle Cycle = iota
	CycleMonth
	CycleYear
)

var cycleNames = [...]string{
	"",
	"month",
	"year",
}

// Chinese translation
var cyclesCN = [...]string{
	"",
	"月",
	"年",
}

// English translation
var cyclesEN = [...]string{
	"",
	"Month",
	"Year",
}

var cycleMap = map[Cycle]string{
	1: cycleNames[1],
	2: cycleNames[2],
}

var cycleValue = map[string]Cycle{
	cycleNames[1]: 1,
	cycleNames[2]: 2,
}

// ParseCycle parses a string into Cycle type.
func ParseCycle(name string) (Cycle, error) {
	if x, ok := cycleValue[name]; ok {
		return x, nil
	}

	return InvalidCycle, fmt.Errorf("%s is not a valid Cycle", name)
}

// TimeAfterACycle adds one cycle to a time instance and returns the new time.
func (c Cycle) TimeAfterACycle(t time.Time) (time.Time, error) {
	switch c {
	case CycleYear:
		return t.AddDate(1, 0, 1), nil
	case CycleMonth:
		return t.AddDate(0, 1, 1), nil
	default:
		return t, errors.New("not a valid cycle type")
	}
}

func (c Cycle) String() string {
	if s, ok := cycleMap[c]; ok {
		return s
	}

	return ""
}

// StringCN output cycle as Chinese text
func (c Cycle) StringCN() string {
	if c < CycleMonth || c > CycleYear {
		return ""
	}

	return cyclesCN[c]
}

// StringEN output cycle as English text
func (c Cycle) StringEN() string {
	if c < CycleMonth || c > CycleYear {
		return ""
	}

	return cyclesEN[c]
}

// UnmarshalJSON implements the Unmarshaler interface.
func (c *Cycle) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tmp, _ := ParseCycle(s)

	*c = tmp

	return nil
}

// MarshalJSON implements the Marshaler interface
func (c Cycle) MarshalJSON() ([]byte, error) {
	s := c.String()
	if s == "" {
		return []byte("null"), nil
	}

	return []byte(`"` + s + `"`), nil
}

// Scan implements sql.Scanner interface to retrieve value from SQL.
// SQL null will be turned into zero value CycleInvalid
func (c *Cycle) Scan(src interface{}) error {
	if src == nil {
		*c = InvalidCycle
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp, _ := ParseCycle(string(s))
		*c = tmp
		return nil

	default:
		return ErrIncompatible
	}
}

// Value implements driver.Valuer interface to save value into SQL.
func (c Cycle) Value() (driver.Value, error) {
	s := c.String()
	if s == "" {
		return nil, nil
	}

	return s, nil
}
