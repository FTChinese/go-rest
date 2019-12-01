package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Platform is used to record on which platform user is visiting the API.
type Platform int

// Allowed values for ClientPlatforms
const (
	PlatformNull Platform = iota
	PlatformWeb
	PlatformIOS
	PlatformAndroid
)

var platformNames = [...]string{
	"",
	"web",
	"ios",
	"android",
}

var platformMap = map[Platform]string{
	1: platformNames[1],
	2: platformNames[2],
	3: platformNames[3],
}

var platformValue = map[string]Platform{
	platformNames[1]: 1,
	platformNames[2]: 2,
	platformNames[3]: 3,
}

// ParsePlatform parses a string into a Platform value.
func ParsePlatform(name string) (Platform, error) {
	if x, ok := platformValue[name]; ok {
		return x, nil
	}

	return PlatformNull, fmt.Errorf("%s is not valid Platform", name)
}

func (x Platform) String() string {
	if s, ok := platformMap[x]; ok {
		return s
	}

	return ""
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *Platform) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tmp, _ := ParsePlatform(s)

	*x = tmp

	return nil
}

// MarshalJSON implements the Marshaler interface
func (x Platform) MarshalJSON() ([]byte, error) {
	s := x.String()

	if s == "" {
		return []byte("null"), nil
	}

	return []byte(`"` + s + `"`), nil
}

// Scan implements sql.Scanner interface to retrieve value from SQL.
// SQL null will be turned into InvalidPlatform.
func (x *Platform) Scan(src interface{}) error {
	if src == nil {
		*x = PlatformNull
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp, _ := ParsePlatform(string(s))
		*x = tmp
		return nil

	default:
		return ErrIncompatible
	}
}

// Value saves ClientPlatform to SQL ENUM.
func (x Platform) Value() (driver.Value, error) {

	s := x.String()
	if s == "" {
		return nil, nil
	}

	return s, nil
}
