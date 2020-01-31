package appErrors

import "errors"

var  (
	GetRequestParamsError = errors.New("Params Error")

	AssertError = errors.New("Assert Error")
)