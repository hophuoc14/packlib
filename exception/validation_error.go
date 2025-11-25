package exception

type ValidationError struct {
	Message string `json:"message"`
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}