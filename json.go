package gorest

import (
	"bytes"
	"encoding/json"
	"io"
)

// ParseJSON parses input data to struct
func ParseJSON(data io.ReadCloser, v interface{}) error {
	dec := json.NewDecoder(data)
	defer data.Close()

	return dec.Decode(v)
}

// Stringify turn an interface into json string
func Stringify(v interface{}) ([]byte, error) {
	// bytes.Buffer implements io.Writer
	buf := new(bytes.Buffer)
	// NewEncoder accepts an io.Writer
	enc := json.NewEncoder(buf)
	// Do not escape HTML
	enc.SetEscapeHTML(false)
	// Indent with a tab.
	// Is it necessary? Browser has tool to format JSON output properly.
	// Firefox format JSON natively.
	// Chrome could use `json` plugin:
	// https://chrome.google.com/webstore/detail/jsonview/chklaanhfefbnpoihckbnefhakgolnmc?hl=en
	enc.SetIndent("", "\t")

	err := enc.Encode(v)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
