package rand

// IntRange returns, as an int, a non-negative pseudo-random number in [min, max).
// It panics if either of min <= 0 or max <= 0.
func IntRange(min, max int) int {
	return seedRand.Intn(max-min) + min
}
