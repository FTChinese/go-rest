package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type SubsStatus int

const (
	SubsStatusNull SubsStatus = iota
	SubStatusActive
	SubsStatusCanceled
	SubStatusIncomplete
	SubsStatusIncompleteExpired
	SubsStatusPastDue
	SubsStatusTrialing
	SubsStatusUnpaid
)

var subsStatusNames = [...]string{
	"",
	"active",
	"canceled",
	"incomplete",
	"incomplete_expired",
	"past_due",
	"trialing",
	"unpaid",
}

// Map SubsStatus to string value to be persisted.
var subsStatusMap = map[SubsStatus]string{
	1: subsStatusNames[1],
	2: subsStatusNames[2],
	3: subsStatusNames[3],
	4: subsStatusNames[4],
	5: subsStatusNames[5],
	6: subsStatusNames[6],
	7: subsStatusNames[7],
}

// Parse a string to SubsStatus
var subsStatusValue = map[string]SubsStatus{
	subsStatusNames[1]: 1,
	subsStatusNames[2]: 2,
	subsStatusNames[3]: 3,
	subsStatusNames[4]: 4,
	subsStatusNames[5]: 5,
	subsStatusNames[6]: 6,
	subsStatusNames[7]: 7,
}

// ParseSubsStatus turns a string to SubsStatus.
func ParseSubsStatus(name string) (SubsStatus, error) {
	if x, ok := subsStatusValue[name]; ok {
		return x, nil
	}

	return SubsStatusNull, fmt.Errorf("%s is not valid SubsStatus", name)
}

// ShouldCreate checks whether membership's current status
// should allow creation of a new membership.
func (x SubsStatus) ShouldCreate() bool {
	return x == SubsStatusNull ||
		x == SubsStatusIncompleteExpired ||
		x == SubsStatusPastDue ||
		x == SubsStatusCanceled ||
		x == SubsStatusUnpaid
}

func (x SubsStatus) String() string {
	if s, ok := subsStatusMap[x]; ok {
		return s
	}

	return ""
}

func (x *SubsStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tmp, _ := ParseSubsStatus(s)

	*x = tmp

	return nil
}

func (x SubsStatus) MarshalJSON() ([]byte, error) {
	s := x.String()

	if s == "" {
		return []byte("null"), nil
	}

	return []byte(`"` + s + `"`), nil
}

func (x *SubsStatus) Scan(src interface{}) error {
	if src == nil {
		*x = SubsStatusNull
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp, _ := ParseSubsStatus(string(s))
		*x = tmp
		return nil

	default:
		return errors.New("incompatible type to scan")
	}
}

func (x SubsStatus) Value() (driver.Value, error) {
	s := x.String()
	if s == "" {
		return nil, nil
	}

	return s, nil
}
