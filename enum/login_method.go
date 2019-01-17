package enum

import (
	"database/sql/driver"
	"fmt"
)

// LoginMethod is an enumeration of login method.
type LoginMethod int

// Allowed values for LoginMethod
const (
	InvalidLogin LoginMethod = iota
	LoginMethodEmail
	LoginMethodWx
)

var loginMethodNames = [...]string{
	"",
	"email",
	"wechat",
}

var loginMethodMap = map[LoginMethod]string{
	1: loginMethodNames[1],
	2: loginMethodNames[2],
}

var loginMethodValue = map[string]LoginMethod{
	loginMethodNames[1]: 1,
	loginMethodNames[2]: 2,
}

// ParseLoginMethod creates a new LoginMethod from a string: email or wechat.
func ParseLoginMethod(name string) (LoginMethod, error) {
	if x, ok := loginMethodValue[name]; ok {
		return x, nil
	}

	return InvalidLogin, fmt.Errorf("%s is not a valid LoginMethod", name)
}

func (x LoginMethod) String() string {
	if str, ok := loginMethodMap[x]; ok {
		return str
	}

	return ""
}

// Scan implements the Scanner interface
func (x *LoginMethod) Scan(value interface{}) error {
	var name string
	switch v := value.(type) {
	case string:
		name = v
	case []byte:
		name = string(v)
	case nil:
		*x = InvalidLogin
		return nil
	}

	tmp, err := ParseLoginMethod(name)

	if err != nil {
		return err
	}

	*x = tmp
	return nil
}

// Value implements the Valuer interface.
func (x LoginMethod) Value() (driver.Value, error) {
	if x == InvalidLogin {
		return nil, nil
	}

	return x.String(), nil
}
