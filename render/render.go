package render

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// Render wraps http.ResponseWrtier and configuration of
// how to send reponse content.
type Render struct {
	writer     http.ResponseWriter
	escapeHTML bool
	indent     string
}

// New createa a new instance of Render.
// By default, HTML in the content is not escaped,
// and JSON is indented with a tab.
func New(w http.ResponseWriter) *Render {
	return &Render{
		writer:     w,
		escapeHTML: false,
		indent:     "\t",
	}
}

// EscapeHTML set escaping HTML.
func (r *Render) EscapeHTML() *Render {
	r.escapeHTML = true
	return r
}

// NoCache set headers to avoid caching repsonse.
func (r *Render) NoCache() *Render {
	r.writer.Header().Add("Cache-Control", "no-cache")
	r.writer.Header().Add("Cache-Control", "no-store")
	r.writer.Header().Add("Cache-Control", "must-revalidate")
	r.writer.Header().Add("Pragma", "no-cache")

	return r
}

func (r *Render) HTML(code int, body string) error {
	r.writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	if body == "" || code == http.StatusNoContent {
		r.writer.WriteHeader(code)
		return nil
	}

	r.writer.WriteHeader(code)

	_, err := r.writer.Write([]byte(body))
	if err != nil {
		return err
	}

	return nil
}

func (r *Render) Text(code int, body string) error {
	r.writer.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if body == "" || code == http.StatusNoContent {
		r.writer.WriteHeader(code)
		return nil
	}

	r.writer.WriteHeader(code)

	_, err := r.writer.Write([]byte(body))
	if err != nil {
		return err
	}

	return nil
}

// JSON renders JSON response.
func (r *Render) JSON(code int, body interface{}) error {
	if r.writer.Header().Get("Content-Type") == "" {
		r.writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	}

	if body == nil || code == http.StatusNoContent {
		r.writer.WriteHeader(code)
		return nil
	}

	r.writer.WriteHeader(code)

	enc := json.NewEncoder(r.writer)
	enc.SetEscapeHTML(r.escapeHTML)
	enc.SetIndent("", r.indent)

	return enc.Encode(body)
}

// OK sends 200 OK response for JSON.
func (r *Render) OK(body interface{}) error {
	return r.JSON(http.StatusOK, body)
}

// NoContent send a 204 response.
func (r *Render) NoContent() error {
	return r.JSON(http.StatusNoContent, nil)
}

// HandleError sends response above 400
func (r *Render) HandleError(re *ResponseError) error {
	re = &ResponseError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
	}

	return r.JSON(re.StatusCode, re)
}

// NotFound sends 404 Not Found response.
func (r *Render) NotFound(msg string) error {
	if msg == "" {
		msg = "Not Found"
	}
	return r.JSON(http.StatusNotFound, ResponseError{
		Message: "Not Found",
	})
}

// Unauthorized sends 401 Unauthorized response.
func (r *Render) Unauthorized(msg string) error {
	if msg == "" {
		msg = "Requires authorization."
	}

	return r.JSON(http.StatusUnauthorized, ResponseError{
		Message: msg,
	})
}

// Forbidden sends 403 response.
func (r *Render) Forbidden(msg string) error {
	if msg == "" {
		msg = "Fobbidden"
	}

	return r.JSON(http.StatusForbidden, ResponseError{
		Message: msg,
	})
}

// BadRequest sends 400 reponse.
func (r *Render) BadRequest(msg string) error {
	return r.JSON(http.StatusBadRequest, ResponseError{
		Message: msg,
	})
}

// Unprocessable sends 422 response
func (r *Render) Unprocessable(ve *ValidationError) error {
	return r.JSON(http.StatusUnprocessableEntity, ResponseError{
		Message: ve.Message,
		Invalid: ve,
	})
}

// TooManyRequests sends 429 response.
func (r *Render) TooManyRequests(msg string) error {
	return r.JSON(http.StatusTooManyRequests, ResponseError{
		Message: msg,
	})
}

// InternalServerError sends 500 response.
func (r *Render) InternalServerError(msg string) error {
	return r.JSON(http.StatusInternalServerError, ResponseError{
		Message: msg,
	})
}

// DBError sends 404 or 500 response.
func (r *Render) DBError(err error) error {
	switch err {
	case sql.ErrNoRows:
		return r.NotFound("")

	default:
		return r.InternalServerError(err.Error())
	}
}
