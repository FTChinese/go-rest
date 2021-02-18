package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// OrderKind describes what kind of subscription order
// a user is purchasing.
type OrderKind int

const (
	OrderKindNull OrderKind = iota
	OrderKindCreate
	OrderKindRenew
	OrderKindUpgrade
	OrderKindDowngrade
	OrderKindAddOn
	OrderKindSwitchCycle
)

var orderKindNames = [...]string{
	"",
	"create",
	"renew",
	"upgrade",
	"downgrade",
	"add_on",
	"switch_cycle", // This is not persisted to db.
}

// String representation of OrderKind
var orderKindMap = map[OrderKind]string{
	1: orderKindNames[1],
	2: orderKindNames[2],
	3: orderKindNames[3],
	4: orderKindNames[4],
	5: orderKindNames[5],
	6: orderKindNames[6],
}

// Simplified Chinese version of OrderKind's string representation.
var orderKindSCMap = map[OrderKind]string{
	1: "订阅",
	2: "续订",
	3: "升级订阅",
	4: "购买标准版",
	5: "补充包",
	6: "更改订阅周期",
}

// Used to get OrderKind from a string.
var orderKindValue = map[string]OrderKind{
	orderKindNames[1]: 1,
	orderKindNames[2]: 2,
	orderKindNames[3]: 3,
	orderKindNames[4]: 4,
	orderKindNames[5]: 5,
	orderKindNames[6]: 6,
}

// ParseOrderKind creates OrderKind from a string.
func ParseOrderKind(name string) (OrderKind, error) {
	if x, ok := orderKindValue[name]; ok {
		return x, nil
	}

	return OrderKindNull, fmt.Errorf("%s is not valid OrderKind", name)
}

func (x OrderKind) String() string {
	if s, ok := orderKindMap[x]; ok {
		return s
	}

	return ""
}

func (x OrderKind) StringSC() string {
	if s, ok := orderKindSCMap[x]; ok {
		return s
	}

	return ""
}

func (x *OrderKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tmp, _ := ParseOrderKind(s)

	*x = tmp

	return nil
}

func (x OrderKind) MarshalJSON() ([]byte, error) {
	s := x.String()

	if s == "" {
		return []byte("null"), nil
	}

	return []byte(`"` + s + `"`), nil
}

func (x *OrderKind) Scan(src interface{}) error {
	if src == nil {
		*x = OrderKindNull
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp, _ := ParseOrderKind(string(s))
		*x = tmp
		return nil

	default:
		return errors.New("incompatible type to scan")
	}
}

func (x OrderKind) Value() (driver.Value, error) {
	s := x.String()
	if s == "" {
		return nil, nil
	}

	return s, nil
}
