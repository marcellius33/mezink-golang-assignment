package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Payload interface{} `json:"records,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func WriteJsonResponse(ctx *gin.Context, resp *Response) {
	ctx.JSON(resp.Code, resp)
}

func InternalServerError(err error) *Response {
	return &Response{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Error:   err.Error(),
	}
}

func BadRequestResponse(err error) *Response {
	return &Response{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Error:   err.Error(),
	}
}

func SuccessResponse(payload interface{}) *Response {
	return &Response{
		Code:    http.StatusOK,
		Message: "Success",
		Payload: payload,
	}
}
