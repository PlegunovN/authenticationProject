package users

type ErrorPasswordIncorrect struct{}

func (e ErrorPasswordIncorrect) Error() string {
	return "incorrect password"
}
