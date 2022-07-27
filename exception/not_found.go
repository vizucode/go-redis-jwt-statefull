package exception

type NotFound struct {
	MsgError string
}

func NotFoundError(Error string) *NotFound {
	return &NotFound{
		MsgError: Error,
	}
}

func (e *NotFound) Error() string {
	return e.MsgError
}
