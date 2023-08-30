package errormsg

import "fmt"

type CustomError struct {
	Code int		`json:"code"`
	Message string	`json:"message"`
	Details string `json:"details"`
}

func (e CustomError) Error() string {
    return fmt.Sprintf("message: %s", e.Message)
}

func NewErrorMessage(code int, message string, details string) CustomError {
	return CustomError{
		Code: code,
		Message: message,
		Details: details,
	}
}