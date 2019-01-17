package enum

import (
	"database/sql/driver"
	"fmt"
)

// ClientPlatform is used to record on which platoform user is visiting the API.
type Platform int

// Allowed values for ClientPlatforms
const (
	InvalidPlatform Platform = iota
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

	return InvalidPlatform, fmt.Errorf("%s is not valid Platform", name)
}

func (x Platform) String() string {
	if s, ok := platformMap[x]; ok {
		return s
	}

	return ""
}

// Scan implements sql.Scanner interface to retrieve value from SQL.
// SQL null will be turned into InvalidPlatform.
func (x *Platform) Scan(src interface{}) error {
	if src == nil {
		*x = InvalidPlatform
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp, err := ParsePlatform(string(s))
		if err != nil {
			return err
		}
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
