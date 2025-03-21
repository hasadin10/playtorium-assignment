package frameworks

import (
	"discountmodule/interfaces"
)

type HttpRequest struct{}

func NewHttpReqest() interfaces.HttpRequest {
	return &HttpRequest{}
}
