package dto

type ErrorMessage struct {
	Message string
}

func NewErrorMessage(msg string) *ErrorMessage {
	return &ErrorMessage{
		Message: msg,
	}
}

type ErrorMessageComplete struct {
	Message     string
	Description string
}

func NewErrorMessageComplete(msg, description string) *ErrorMessageComplete {
	return &ErrorMessageComplete{
		Message:     msg,
		Description: description,
	}
}
