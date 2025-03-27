package errors

type Err struct {
	Typ ErrType
	Msg string
	Err error
}

func NewErr(typ ErrType, msg string, err error) *Err {
	return &Err{Typ: typ, Msg: msg, Err: err}
}

func (e Err) Error() string {
	return e.Msg
}
