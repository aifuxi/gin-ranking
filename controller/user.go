package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (uc UserController) Batch(c *gin.Context) {
	ReturnSuccess(c, &Response{
		Code: 0,
		Msg:  "success",
		Data: "user batch",
	})
}

func (uc UserController) Create(c *gin.Context) {
	ReturnSuccess(c, &Response{
		Code: 0,
		Msg:  "success",
		Data: "user create",
	})
}

func (uc UserController) Update(c *gin.Context) {
	ReturnSuccess(c, &Response{
		Code: 0,
		Msg:  "success",
		Data: "user update",
	})
}

func (uc UserController) Delete(c *gin.Context) {
	ReturnSuccess(c, &Response{
		Code: 0,
		Msg:  "success",
		Data: "user delete",
	})
}
