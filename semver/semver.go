package semver

import (
	"strconv"
	"strings"
)

// SemVer represents a parsed value of semantic version string.
type SemVer struct {
	Major int
	Minor int
	Patch int
}

// Parse parsed a string into to SemVer
func Parse(v string) SemVer {
	var parts = make([]int, 0)

	for _, v := range strings.Split(v, ".") {
		n, err := strconv.Atoi(v)
		if err != nil {
			n = 0
		}

		parts = append(parts, n)
	}

	gap := len(parts) - 3
	if gap > 0 {
		for i := 0; i < gap; i++ {
			parts = append(parts, 0)
		}
	}

	return SemVer{
		Major: parts[0],
		Minor: parts[1],
		Patch: parts[2],
	}
}

// Compare compares two semantic versions.
// Returns negative number if s should come before (is smaller) other;
// 0 if the two are equal;
// Positive number if s should come after (is larger than) other.
func (s SemVer) Compare(other SemVer) int {
	diff := s.Major - other.Major
	if diff != 0 {
		return diff
	}

	diff = s.Minor - other.Minor
	if diff != 0 {
		return diff
	}

	return s.Patch - other.Patch
}

// Equal tests if two semantic version string are the same.
func (s SemVer) Equal(other SemVer) bool {
	return s.Compare(other) == 0
}

// Larger tests whether semantic version a is larger than b
func (s SemVer) Larger(other SemVer) bool {
	return s.Compare(other) > 0
}

// Smaller tests whether semantic version a is smaller than b.
func (s SemVer) Smaller(other SemVer) bool {
	return s.Compare(other) < 0
}
