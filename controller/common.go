package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Msg  any `json:"msg"`
	Data any `json:"data"`
}

func ReturnSuccess(c *gin.Context, resp *Response) {

	c.JSON(http.StatusOK, resp)
}
