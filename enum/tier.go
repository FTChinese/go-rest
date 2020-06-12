package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Tier is an enum for membership tiers.
type Tier int

// Values of MemberTier
const (
	TierNull Tier = iota
	TierStandard
	TierPremium
	TierVIP
)

var tierNames = [...]string{
	"",
	"standard",
	"premium",
	"vip",
}

var tiersCN = [...]string{
	"",
	"标准会员",
	"高级会员",
	"VIP",
}

var tiersEN = [...]string{
	"",
	"Standard",
	"Premium",
	"VIP",
}

var tierMap = map[Tier]string{
	1: tierNames[1],
	2: tierNames[2],
	3: tierNames[3],
}

var tierValue = map[string]Tier{
	tierNames[1]: 1,
	tierNames[2]: 2,
	tierNames[3]: 3,
}

// ParseTier parses a string into Tier type.
func ParseTier(name string) (Tier, error) {
	if x, ok := tierValue[name]; ok {
		return x, nil
	}

	return TierNull, fmt.Errorf("%s is not a valid Tier", name)
}

func (x Tier) String() string {
	if s, ok := tierMap[x]; ok {
		return s
	}

	return ""
}

// StringCN output tier as Chinese text
func (x Tier) StringCN() string {
	if x < TierStandard || x > TierPremium {
		return ""
	}

	return tiersCN[x]
}

// StringEN output tier as English text
func (x Tier) StringEN() string {
	if x < TierStandard || x > TierPremium {
		return ""
	}

	return tiersEN[x]
}

// UnmarshalJSON implements the Unmarshaler interface.
func (x *Tier) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tmp, _ := ParseTier(s)

	*x = tmp

	return nil
}

// MarshalJSON implements the Marshaler interface
func (x Tier) MarshalJSON() ([]byte, error) {
	s := x.String()

	if s == "" {
		return []byte("null"), nil
	}

	return []byte(`"` + s + `"`), nil
}

// Scan implements sql.Scanner interface to retrieve value from SQL.
// SQL null will be turned into zero value TierFree.
func (x *Tier) Scan(src interface{}) error {
	if src == nil {
		*x = TierNull
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp, _ := ParseTier(string(s))
		*x = tmp
		return nil

	default:
		return ErrIncompatible
	}
}

// Value implements driver.Valuer interface to save value into SQL.
func (x Tier) Value() (driver.Value, error) {
	s := x.String()
	if s == "" {
		return nil, nil
	}

	return s, nil
}
