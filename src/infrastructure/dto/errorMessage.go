package dto

type ErrorMessage struct {
	Message string
}

func NewErrorMessage(msg string) *ErrorMessage {
	return &ErrorMessage{
		Message: msg,
	}
}
