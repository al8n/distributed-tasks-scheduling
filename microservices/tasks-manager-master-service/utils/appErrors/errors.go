package appErrors

import "errors"

var  (
	GetRequestParamsError = errors.New("Params Error")

	AssertError = errors.New("Assert Error")
	EndpointAssertError = errors.New("Endport Assert Error")
	DecoderAssertError = errors.New("Decoder Assert Error")
	EncoderAssertError = errors.New("Encoder Assert Error")
)