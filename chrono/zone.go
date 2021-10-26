package chrono

import "time"

const (
	secondsOfMinute = 60
	secondsOfHour   = 60 * secondsOfMinute
)

// TZShanghai is fixed time zone set to UTC8.
var (
	TZShanghai = time.FixedZone("UTC+8", 8*secondsOfHour)
)
