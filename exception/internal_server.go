package exception

type InternalServer struct {
	MsgError string
}

func InternalServerError(Error string) *NotFound {
	return &NotFound{
		MsgError: Error,
	}
}

func (e *InternalServer) Error() string {
	return e.MsgError
}
