package chrono

import "time"

const (
	secondsOfMinute = 60
	secondsOfHour   = 60 * secondsOfMinute
)

// Fixed time zones.
var (
	TZShanghai = time.FixedZone("UTC+8", 8*secondsOfHour)
)