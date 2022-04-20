package exceptions

type InvalidCredentialsException struct {
}

func (e *InvalidCredentialsException) Error() string {
	return "Invalid email or password."
}
