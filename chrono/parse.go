package chrono

import (
	"fmt"
	"time"
)

// ParseDateTime parses SQL DATE or DATETIME string in specified location.
func ParseDateTime(str string, loc *time.Location) (t time.Time, err error) {
	base := "0000-00-00 00:00:00.0000000"
	switch len(str) {
	case 10, 19: // up to "YYYY-MM-DD HH:MM:SS"
		if str == base[:len(str)] {
			return
		}
		t, err = time.Parse(SQLDateTime[:len(str)], str)
	default:
		err = fmt.Errorf("invalid time string: %s", str)
		return
	}

	// Adjust location
	if err == nil && loc != time.UTC {
		y, mo, d := t.Date()
		h, mi, s := t.Clock()
		t, err = time.Date(y, mo, d, h, mi, s, t.Nanosecond(), loc), nil
	}

	return
}