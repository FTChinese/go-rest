package gorest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Param represents a pair of query parameter from URL.
type Param struct {
	key   string
	value string
}

// GetQueryParam get a pair of query parameter from URL.
func GetQueryParam(req *http.Request, key string) Param {
	v := req.Form.Get(key)

	return Param{
		key:   key,
		value: strings.TrimSpace(v),
	}
}

// ToBool converts a query parameter to boolean value.
func (p Param) ToBool() (bool, error) {
	return strconv.ParseBool(string(p.value))
}

// ToString converts a query parameter to string value.
// Returns error for an empty value.
func (p Param) ToString() (string, error) {
	if p.value == "" {
		return "", fmt.Errorf("%s have empty value", p.key)
	}

	return p.value, nil
}

// ToInt converts the value of a query parameter to int64
func (p Param) ToInt() (int64, error) {
	if p.value == "" {
		return 0, fmt.Errorf("%s have empty value", p.key)
	}

	num, err := strconv.ParseInt(p.value, 10, 0)

	if err != nil {
		return 0, err
	}

	return num, nil
}
