package service

const (
	FieldEmpty ErrorCode = iota
)

type (
	ErrorCode       int
	ValidationError struct {
		Code    ErrorCode
		Message string
		Field   string
	}
)

func (e ValidationError) Error() string {
	return e.Message
}

func NewValidationError(code ErrorCode, message string, field string) ValidationError {
	return ValidationError{
		Code:    code,
		Message: message,
		Field:   field,
	}
}
