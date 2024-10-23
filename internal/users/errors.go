package users

type ErrorPasswordIncorrect struct{}

func (e ErrorPasswordIncorrect) Error() string {
	return "incorrect password"
}

type ErrorDuplicateLogin struct {
}

func (e ErrorDuplicateLogin) Error() string {
	return "error duplicate login"
}
