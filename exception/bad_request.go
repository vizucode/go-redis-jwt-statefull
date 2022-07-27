package exception

type BadRequest struct {
	MsgError string
}

func BadRequestError(Error string) *BadRequest {
	return &BadRequest{
		MsgError: Error,
	}
}

func (e *BadRequest) Error() string {
	return e.MsgError
}
