package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code  int `json:"code"`
	Msg   any `json:"msg"`
	Data  any `json:"data,omitempty"`
	Error any `json:"error,omitempty"`
	Total int `json:"total,omitempty"`
}

type Option func(resp *Response)

func WithCode(code int) Option {
	return func(resp *Response) {
		resp.Code = code
	}
}

func WithMsg(msg string) Option {
	return func(resp *Response) {
		resp.Msg = msg
	}
}

func WithTotal(total int) Option {
	return func(resp *Response) {
		resp.Total = total
	}
}

func WithData(data any) Option {
	return func(resp *Response) {
		resp.Data = data
	}
}

func WithError(err any) Option {
	return func(resp *Response) {
		resp.Error = err
	}
}

func ReturnResponse(c *gin.Context, opts ...Option) {
	resp := &Response{
		Code:  0,
		Msg:   "success",
		Data:  nil,
		Error: nil,
		Total: 0,
	}

	for _, opt := range opts {
		opt(resp)
	}

	c.JSON(http.StatusOK, *resp)
}
