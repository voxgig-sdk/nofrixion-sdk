package core

type NofrixionError struct {
	IsNofrixionError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewNofrixionError(code string, msg string, ctx *Context) *NofrixionError {
	return &NofrixionError{
		IsNofrixionError: true,
		Sdk:              "Nofrixion",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *NofrixionError) Error() string {
	return e.Msg
}
