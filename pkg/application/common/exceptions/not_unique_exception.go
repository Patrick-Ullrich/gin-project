package exceptions

type NotUniqueException struct {
	Field   string
	Message string
}

func (e *NotUniqueException) Error() string {
	return e.Message
}
