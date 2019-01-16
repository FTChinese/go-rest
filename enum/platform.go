package enum

import (
	"database/sql/driver"
	"fmt"
)

// ClientPlatform is used to record on which platoform user is visiting the API.
type Platform int

const InvalidPlatform Platform = -1

// Allowed values for ClientPlatforms
const (
	PlatformWeb Platform = iota
	PlatformIOS
	PlatformAndroid
)

var platformNames = [...]string{
	"web",
	"ios",
	"android",
}

var platformMap = map[Platform]string{
	0: platformNames[0],
	1: platformNames[1],
	2: platformNames[2],
}

var platformValue = map[string]Platform{
	platformNames[0]: 0,
	platformNames[1]: 1,
	platformNames[2]: 2,
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
