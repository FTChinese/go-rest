package enum

import (
	"database/sql/driver"
	"encoding/json"
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

// UnmarshalJSON implements the Unmarshaler interface.
func (x *LoginMethod) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	// Error should be ignored since any values not allowed should be turned into  InvalidLogin.
	tmp, _ := ParseLoginMethod(s)

	*x = tmp

	return nil
}

// MarshalJSON implements the Marshaler interface
func (x LoginMethod) MarshalJSON() ([]byte, error) {
	s := x.String()

	if s == "" {
		return []byte("null"), nil
	}

	return []byte(`"` + s + `"`), nil
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

	// Ignore error.
	tmp, _ := ParseLoginMethod(name)

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
