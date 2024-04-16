package helpers

type GenericError struct {
	Msg  string
	Code int
	Err  string
}

func (e *GenericError) Error() string {
	if e.Code == 0 {
		e.Code = 400
	}

	return e.Msg
}
