package render

import (
	"database/sql"
	"fmt"
	"net/http"
)

// InvalidCode is the reason why validation failed.
type InvalidCode string

const (
	// CodeMissing means a resource does not exist
	CodeMissing InvalidCode = "missing"
	// CodeMissingField means a required field on a resource has not been set.
	CodeMissingField InvalidCode = "missing_field"
	// CodeInvalid means the formatting of a field is invalid
	CodeInvalid InvalidCode = "invalid"
	// CodeAlreadyExists means another resource has the same value as this field.
	CodeAlreadyExists InvalidCode = "already_exists"
)

// ValidationError tells the field that failed validation.
type ValidationError struct {
	Message string      `json:"-"`
	Field   string      `json:"field"`
	Code    InvalidCode `json:"code"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

// ResponseError is the response body for http code above 400.
type ResponseError struct {
	StatusCode int              `json:"-"`
	Message    string           `json:"message"`
	Invalid    *ValidationError `json:"error,omitempty"`
}

func (re *ResponseError) Error() string {
	return fmt.Sprintf("code=%d, message=%s", re.StatusCode, re.Message)
}

// NewResponseError creates a new ResponseError instance.
func NewResponseError(code int, msg string) *ResponseError {
	return &ResponseError{
		StatusCode: code,
		Message:    msg,
	}
}

// NewNotFound creates response 404 Not Found
func NewNotFound(msg string) *ResponseError {
	return NewResponseError(http.StatusNotFound, msg)
}

// NewUnauthorized create a new instance of Response for 401 Unauthorized response
func NewUnauthorized(msg string) *ResponseError {
	if msg == "" {
		msg = "Requires authorization."
	}

	return NewResponseError(http.StatusUnauthorized, msg)
}

// NewForbidden creates response for 403
func NewForbidden(msg string) *ResponseError {
	return NewResponseError(http.StatusForbidden, msg)
}

// NewBadRequest creates a new Response for 400 Bad Request with the specified msg
func NewBadRequest(msg string) *ResponseError {
	return NewResponseError(http.StatusBadRequest, msg)
}

// NewUnprocessable creates response 422 Unprocessable Entity
func NewUnprocessable(ve *ValidationError) *ResponseError {

	return &ResponseError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    ve.Message,
		Invalid:    ve,
	}
}

// NewAlreadyExists is a convenience func to handle MySQL
// 1062 error.
func NewAlreadyExists(field string) *ResponseError {
	return NewUnprocessable(&ValidationError{
		Message: "Duplicate entry",
		Field:   field,
		Code:    CodeAlreadyExists,
	})
}

// NewTooManyRequests respond to rate limit.
func NewTooManyRequests(msg string) *ResponseError {
	return NewResponseError(http.StatusTooManyRequests, msg)
}

// NewInternalError creates response for internal server error
func NewInternalError(msg string) *ResponseError {

	return NewResponseError(http.StatusInternalServerError, msg)
}

// NewDBError handles various errors returned from the model layer
// MySQL duplicate error when inserting into uniquely constraint column;
// ErrNoRows if it cannot retrieve any rows of the specified criteria;
// `field` is used to identify which field is causing duplicate error.
func NewDBError(err error) *ResponseError {
	switch err {
	case sql.ErrNoRows:
		return NewNotFound("")

	default:
		return NewInternalError(err.Error())
	}
}
