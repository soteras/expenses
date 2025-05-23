package errors

type AuthError struct {
	Msg string
}

func (e *AuthError) Error() string {
	return e.Msg
}

func NewAuthError(msg string) error {
	return &AuthError{Msg: msg}
}
