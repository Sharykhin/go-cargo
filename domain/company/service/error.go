package service

const (
	Empty = 0
)

type (
	ValidationError struct {
		Code int
		Message string
	}
)

func (e ValidationError) Error() string {
	return e.Message
}

func NewValidationError(code int, message string) ValidationError {
	return ValidationError{
		Code: code,
		Message: message,
	}
}