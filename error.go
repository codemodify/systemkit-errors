package errors

import "fmt"

const (
	ErrCodeNone = -1
)

// Error - defines a detailed error model
type Error struct {
	Occur   bool   `json:"occur"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// NewErrorWithCode - returns `Error` object
func NewErrorWithCode(code int, message string) Error {
	return Error{
		Occur:   true,
		Code:    code,
		Message: message,
	}
}

// NewError - convenience shortcut to `NewErrorWithCode`
func NewError(message string) Error {
	return NewErrorWithCode(ErrCodeNone, message)
}

// New - returns classic `error`
func New(message string) error {
	return NewError(message)
}

// String - implements `Stringer interface`
func (thisRef Error) String() string {
	return fmt.Sprintf("occur: %t, code: %d, message: %s", thisRef.Occur, thisRef.Code, thisRef.Message)
}

// Error - implements `error interface`
func (thisRef Error) Error() string {
	return thisRef.String()
}
