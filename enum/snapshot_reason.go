package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// SnapshotReason tells why we take a snapshot of reader's membership
type SnapshotReason int

const (
	SnapshotReasonNull SnapshotReason = iota
	SnapshotReasonRenew
	SnapshotReasonUpgrade
	SnapshotReasonDelete
	SnapshotReasonLink
	SnapshotReasonUnlink
	SnapshotReasonAppleIAP
)

var snapshotReasonNames = [...]string{
	"",
	"renew",
	"upgrade",
	"delete",
	"link",
	"unlink",
	"apple_iap",
}

var snapshotReasonMap = map[SnapshotReason]string{
	SnapshotReasonRenew:    snapshotReasonNames[1],
	SnapshotReasonUpgrade:  snapshotReasonNames[2],
	SnapshotReasonDelete:   snapshotReasonNames[3],
	SnapshotReasonLink:     snapshotReasonNames[4],
	SnapshotReasonUnlink:   snapshotReasonNames[5],
	SnapshotReasonAppleIAP: snapshotReasonNames[6],
}

var snapshotReasonValue = map[string]SnapshotReason{
	snapshotReasonNames[1]: SnapshotReasonRenew,
	snapshotReasonNames[2]: SnapshotReasonUpgrade,
	snapshotReasonNames[3]: SnapshotReasonDelete,
	snapshotReasonNames[4]: SnapshotReasonLink,
	snapshotReasonNames[5]: SnapshotReasonUnlink,
	snapshotReasonNames[6]: SnapshotReasonAppleIAP,
}

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
