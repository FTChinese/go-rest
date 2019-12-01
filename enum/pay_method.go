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
	PayMethodNull PayMethod = iota
	PayMethodAli
	PayMethodWx
	PayMethodStripe
	PayMethodApple
)

var payMethodNames = [...]string{
	"",
	"alipay",
	"wechat",
	"stripe",
	"apple",
}

var payMethodCN = [...]string{
	"",
	"支付宝",
	"微信支付",
	"Stripe",
	"Apple内购",
}

var payMethodEN = [...]string{
	"",
	"Alipay",
	"Wechat Pay",
	"Stripe",
	"Apple IAP",
}

var payMethodMap = map[PayMethod]string{
	PayMethodAli:    payMethodNames[1],
	PayMethodWx:     payMethodNames[2],
	PayMethodStripe: payMethodNames[3],
	PayMethodApple:  payMethodNames[4],
}

var payMethodValue = map[string]PayMethod{
	payMethodNames[1]: PayMethodAli,
	payMethodNames[2]: PayMethodWx,
	payMethodNames[3]: PayMethodStripe,
	payMethodNames[4]: PayMethodApple,
	"tenpay":          PayMethodWx,
}

// ParsePayMethod parses a string into a PayMethod value.
func ParsePayMethod(name string) (PayMethod, error) {
	if x, ok := payMethodValue[name]; ok {
		return x, nil
	}

	return PayMethodNull, fmt.Errorf("%s is not a valid PayMethod", name)
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

// StringEN output cycle as English text
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

	tmp, _ := ParsePayMethod(s)

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
		*x = PayMethodNull
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp, _ := ParsePayMethod(string(s))
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
