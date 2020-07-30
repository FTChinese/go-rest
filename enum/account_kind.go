package enum

import (
	"encoding/json"
	"fmt"
)

// AccountKind is an enumeration of login method.
type AccountKind int

// Allowed values for AccountKind
const (
	AccountKindNull AccountKind = iota
	AccountKindFtc
	AccountKindWx
	AccountKindLinked
)

var accountKindNames = [...]string{
	"",
	"ftc",
	"wechat",
	"linked",
}

var accountKindMap = map[AccountKind]string{
	1: accountKindNames[1],
	2: accountKindNames[2],
	3: accountKindNames[3],
}

var accountKindValue = map[string]AccountKind{
	accountKindNames[1]: AccountKindFtc,
	accountKindNames[2]: AccountKindWx,
	accountKindNames[3]: AccountKindLinked,
}

func ParseAccountKind(name string) (AccountKind, error) {
	if x, ok := accountKindValue[name]; ok {
		return x, nil
	}

	return AccountKindNull, fmt.Errorf("%s is not valid AccountKind", name)
}

func (x AccountKind) String() string {
	if str, ok := accountKindMap[x]; ok {
		return str
	}

	return ""
}

func (x *AccountKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tmp, _ := ParseAccountKind(s)

	*x = tmp

	return nil
}

// MarshalJSON implements the Marshaler interface
func (x AccountKind) MarshalJSON() ([]byte, error) {
	s := x.String()

	if s == "" {
		return []byte("null"), nil
	}

	return []byte(`"` + s + `"`), nil
}
