package MyError

import "fmt"


type OpError struct {
	code int
	msg string
}

func (p *OpError) Error() string {
	return fmt.Sprintf("Error : %d, %s", p.code, p.msg)
}

func (p *OpError) GetCode() int {
	return p.code
}

func (p *OpError) GetMsg() string {
	return p.msg
}

func NewError(code int, msg string) *OpError{
	return &OpError {
		code:code,
		msg: msg,
	}
}

