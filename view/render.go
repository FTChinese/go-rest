package view

import (
	"encoding/json"
	"net/http"
)

// Render responds to client request
func Render(w http.ResponseWriter, resp Response) error {
	// Set response headers
	for key, vals := range resp.Header {
		for _, v := range vals {
			w.Header().Add(key, v)
		}
	}

	// If `Content-Type` is not set, set the json
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	}

	// If there's no content body, or status code is 204, stop here.
	if resp.Body == nil || resp.StatusCode == http.StatusNoContent {
		w.WriteHeader(resp.StatusCode)
		return nil
	}

	w.WriteHeader(resp.StatusCode)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	enc.SetIndent("", "\t")

	// Write data to w
	return enc.Encode(resp.Body)
}
