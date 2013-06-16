package models

type Error struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

func BuildError(code string, message string) (error *Error) {
	return &Error{
		code,
		message,
	}
}
