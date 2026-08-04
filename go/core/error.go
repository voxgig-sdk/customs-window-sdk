package core

type CustomsWindowError struct {
	IsCustomsWindowError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewCustomsWindowError(code string, msg string, ctx *Context) *CustomsWindowError {
	return &CustomsWindowError{
		IsCustomsWindowError: true,
		Sdk:              "CustomsWindow",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *CustomsWindowError) Error() string {
	return e.Msg
}
