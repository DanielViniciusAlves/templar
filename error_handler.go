package templar

import "fmt"

type CustomError struct {
	Message      string
	DefaultError error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error in Templar: %s, %v.", e.Message, e.DefaultError)
}
