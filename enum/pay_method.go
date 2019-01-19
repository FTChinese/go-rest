package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// PayMethod is an enum for payment methods
type PayMethod int

// Supported payment methods
const (
	InvalidPay PayMethod = iota
	PayMethodAli
	PayMethodWx
	PayMethodStripe
)

var payMethodNames = [...]string{
	"",
	"alipay",
	"tenpay",
	"stripe",
}

var payMethodCN = [...]string{
	"",
	"支付宝",
	"微信支付",
	"PayMethodStripe",
}

var payMethodEN = [...]string{
	"",
	"Zhifubao",
	"Wechat Pay",
	"PayMethodStripe",
}

var payMethodMap = map[PayMethod]string{
	1: payMethodNames[1],
	2: payMethodNames[2],
	3: payMethodNames[3],
}

var payMethodValue = map[string]PayMethod{
	payMethodNames[1]: 1,
	payMethodNames[2]: 2,
	payMethodNames[3]: 3,
}

// ParsePayMethod parses a string into a PayMethod value.
func ParsePayMethod(name string) (PayMethod, error) {
	if x, ok := payMethodValue[name]; ok {
		return x, nil
	}

	return InvalidPay, fmt.Errorf("%s is not a valid PayMethod", name)
}

func (x PayMethod) String() string {
	if str, ok := payMethodMap[x]; ok {
		return str
	}

	return ""
}

// StringCN output cycle as Chinese text
func (x PayMethod) StringCN() string {
	if x < PayMethodAli || x > PayMethodStripe {
		return ""
	}

	return payMethodCN[x]
}

// StringEn output cycle as English text
func (x PayMethod) StringEN() string {
	if x < PayMethodAli || x > PayMethodStripe {
		return ""
	}

	return payMethodEN[x]
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *PayMethod) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tmp, err := ParsePayMethod(s)

	if err != nil {
		return err
	}

	*x = tmp

	return nil
}

// MarshalJSON implements the Marshaler interface
func (x PayMethod) MarshalJSON() ([]byte, error) {
	str := x.String()

	if str == "" {
		return []byte("null"), nil
	}

	return []byte(`"` + str + `"`), nil
}

// Scan implements sql.Scanner interface to retrieve value from SQL.
// SQL null will be turned into zero value InvalidPay.
func (x *PayMethod) Scan(src interface{}) error {
	if src == nil {
		*x = InvalidPay
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp, err := ParsePayMethod(string(s))
		if err != nil {
			return err
		}
		*x = tmp
		return nil

	default:
		return ErrIncompatible
	}
}

// Value implements driver.Valuer interface to save value into SQL.
func (x PayMethod) Value() (driver.Value, error) {
	s := x.String()
	if s == "" {
		return nil, nil
	}

	return s, nil
}
