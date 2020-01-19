package view

const (
	// CodeMissing means a resource does not exist
	CodeMissing = "missing"
	// CodeMissingField means a required field on a resource has not been set.
	CodeMissingField = "missing_field"
	// CodeInvalid means the formatting of a field is invalid
	CodeInvalid = "invalid"
	// CodeAlreadyExists means another resource has the same value as this field.
	CodeAlreadyExists = "already_exists"
)

// ClientError respond to 4xx http status.
type ClientError struct {
	Message string  `json:"message"`
	Reason  *Reason `json:"error,omitempty"`
	// Integrate Stripe errors
	Code  string `json:"code,omitempty"`  // For some errors that could be handled programmatically, a short string indicating the error code reported.
	Param string `json:"param,omitempty"` // If the error is parameter-specific, the parameter related to the error. For example, you can use this to display a message near the correct form field.
	Type  string `json:"type,omitempty"`  // The type of error returned. One of api_connection_error, api_error, authentication_error, card_error, idempotency_error, invalid_request_error, or rate_limit_error
}

// Reason tells why its unprocessable.
// Mostly used for validation errors.
type Reason struct {
	message string // Deprecated
	Message string `json:"-"`
	Field   string `json:"field"`
	Code    string `json:"code"`
}

// NewReason creates a new instance of Reason
func NewReason() *Reason {
	return &Reason{message: "Validation failed"}
}

// NewInvalid creates a new instance of invalid reason.
func NewInvalid(m string) *Reason {
	return &Reason{message: m}
}

// SetMessage set the message to be carried away.
func (r *Reason) SetMessage(msg string) {
	r.message = msg
}

// GetMessage returns Reason's descriptive message.
func (r *Reason) GetMessage() string {
	return r.message
}
