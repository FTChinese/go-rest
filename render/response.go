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

// InvalidAlreadyExists creates a ValidationError for
// MySQL unique key constraint failure.
func InvalidAlreadyExists(field string) *ValidationError {
	return &ValidationError{
		Message: "Duplicate entry",
		Field:   field,
		Code:    CodeAlreadyExists,
	}
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

// ErrorNotFound creates response 404 Not Found
func ErrorNotFound(msg string) *ResponseError {
	return NewResponseError(http.StatusNotFound, msg)
}

// ErrorUnauthorized create a new instance of Response for 401 Unauthorized response
func ErrorUnauthorized(msg string) *ResponseError {
	if msg == "" {
		msg = "Requires authorization."
	}

	return NewResponseError(http.StatusUnauthorized, msg)
}

// ErrorForbidden creates response for 403
func ErrorForbidden(msg string) *ResponseError {
	return NewResponseError(http.StatusForbidden, msg)
}

// NewBadRequest creates a new Response for 400 Bad Request with the specified msg
func NewBadRequest(msg string) *ResponseError {
	return NewResponseError(http.StatusBadRequest, msg)
}

// ErrorUnprocessable creates response 422 Unprocessable Entity
func ErrorUnprocessable(ve *ValidationError) *ResponseError {

	return &ResponseError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    ve.Message,
		Invalid:    ve,
	}
}

// ErrorAlreadyExists is a convenience func to handle MySQL
// 1062 error.
func ErrorAlreadyExists(field string) *ResponseError {
	return ErrorUnprocessable(InvalidAlreadyExists(field))
}

// ErrorTooManyRequests respond to rate limit.
func ErrorTooManyRequests(msg string) *ResponseError {
	return NewResponseError(http.StatusTooManyRequests, msg)
}

// NewInternalError creates response for internal server error
func NewInternalError(msg string) *ResponseError {

	return NewResponseError(http.StatusInternalServerError, msg)
}

// ErrorDB handles various errors returned from the model layer
// MySQL duplicate error when inserting into uniquely constraint column;
// ErrNoRows if it cannot retrieve any rows of the specified criteria;
// `field` is used to identify which field is causing duplicate error.
func ErrorDB(err error) *ResponseError {
	switch err {
	case sql.ErrNoRows:
		return ErrorNotFound("")

	default:
		return NewInternalError(err.Error())
	}
}
