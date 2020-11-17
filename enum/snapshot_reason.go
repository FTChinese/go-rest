package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// SnapshotReason tells why we take a snapshot of reader's membership
// Deprecated
type SnapshotReason int

// Enum of SnapshotReason.
const (
	SnapshotReasonNull        SnapshotReason = iota // Unknow reason.
	SnapshotReasonRenew                             // Backup before renewal
	SnapshotReasonUpgrade                           // Backup before upgrading.
	SnapshotReasonDelete                            // Backup before deletion.
	SnapshotReasonLink                              // Link FTC account to wechat account.
	SnapshotReasonUnlink                            // Unlink FTC account from wechat accout.
	SnapshotReasonAppleLink                         // Link FTC account to Apple IAP.
	SnapshotReasonAppleUnlink                       // Unlink FTC account from Apple IAP.
	SnapshotReasonB2B
	SnapshotReasonManual
	SnapshotReasonIapUpdate
)

var snapshotReasonNames = [...]string{
	"",
	"renew",
	"upgrade",
	"delete",
	"link",
	"unlink",
	"apple_link",
	"apple_unlink",
	"b2b",
	"manual",
	"iap_update",
}

var snapshotReasonMap = map[SnapshotReason]string{
	SnapshotReasonRenew:       snapshotReasonNames[1],
	SnapshotReasonUpgrade:     snapshotReasonNames[2],
	SnapshotReasonDelete:      snapshotReasonNames[3],
	SnapshotReasonLink:        snapshotReasonNames[4],
	SnapshotReasonUnlink:      snapshotReasonNames[5],
	SnapshotReasonAppleLink:   snapshotReasonNames[6],
	SnapshotReasonAppleUnlink: snapshotReasonNames[7],
	SnapshotReasonB2B:         snapshotReasonNames[8],
	SnapshotReasonManual:      snapshotReasonNames[9],
	SnapshotReasonIapUpdate:   snapshotReasonNames[10],
}

var snapshotReasonValue = map[string]SnapshotReason{
	snapshotReasonNames[1]:  SnapshotReasonRenew,
	snapshotReasonNames[2]:  SnapshotReasonUpgrade,
	snapshotReasonNames[3]:  SnapshotReasonDelete,
	snapshotReasonNames[4]:  SnapshotReasonLink,
	snapshotReasonNames[5]:  SnapshotReasonUnlink,
	snapshotReasonNames[6]:  SnapshotReasonAppleLink,
	snapshotReasonNames[7]:  SnapshotReasonAppleUnlink,
	snapshotReasonNames[8]:  SnapshotReasonB2B,
	snapshotReasonNames[9]:  SnapshotReasonManual,
	snapshotReasonNames[10]: SnapshotReasonIapUpdate,
}

// ParseSnapshotReason turns a string to an instance of SnapshotReason
func ParseSnapshotReason(name string) (SnapshotReason, error) {
	if x, ok := snapshotReasonValue[name]; ok {
		return x, nil
	}

	return SnapshotReasonNull, fmt.Errorf("%s is not valid SnapshotReason", name)
}

func (x SnapshotReason) String() string {
	if s, ok := snapshotReasonMap[x]; ok {
		return s
	}

	return ""
}

func (x *SnapshotReason) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	tmp, _ := ParseSnapshotReason(s)

	*x = tmp

	return nil
}

func (x SnapshotReason) MarshalJSON() ([]byte, error) {
	s := x.String()

	if s == "" {
		return []byte("null"), nil
	}

	return []byte(`"` + s + `"`), nil
}

func (x *SnapshotReason) Scan(src interface{}) error {
	if src == nil {
		*x = SnapshotReasonNull
		return nil
	}

	switch s := src.(type) {
	case []byte:
		tmp, _ := ParseSnapshotReason(string(s))
		*x = tmp
		return nil

	default:
		return errors.New("incompatible type to scan")
	}
}

func (x SnapshotReason) Value() (driver.Value, error) {
	s := x.String()

	if s == "" {
		return nil, nil
	}

	return s, nil
}
