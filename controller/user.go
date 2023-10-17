package controller

import (
	"github.com/aifuxi/gin-ranking/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
}

func (uc UserController) Batch(c *gin.Context) {
	users, err := models.GetUsers()
	if err != nil {
		ReturnResponse(c, WithCode(-1), WithError(err))
		return
	}

	ReturnResponse(c, WithData(users))
}

type createUserParams struct {
	Name     string
	Password string
}

func (uc UserController) Create(c *gin.Context) {
	var params createUserParams

	err := c.ShouldBindJSON(&params)
	if err != nil {
		ReturnResponse(c, WithCode(-1), WithError(err))
		return
	}

	user, err := models.CreateUser(params.Name, params.Password)
	if err != nil {
		ReturnResponse(c, WithCode(-1), WithError(err))
		return
	}

	ReturnResponse(c, WithData(user))
}

type updateUserParams struct {
	Name     string
	Password string
}

func (uc UserController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var params updateUserParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		ReturnResponse(c, WithCode(-1), WithError(err))
		return
	}

	user, err := models.UpdateUser(id, params.Name, params.Password)
	if err != nil {
		ReturnResponse(c, WithCode(-1), WithError(err))
		return
	}

	if user.ID == 0 {
		ReturnResponse(c, WithCode(4004), WithMsg("user not found"))
		return
	}

	ReturnResponse(c, WithData(user))
}

func (uc UserController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := models.DeleteUser(id)
	if err != nil {
		ReturnResponse(c, WithCode(-1), WithError(err))
		return
	}

	ReturnResponse(c, WithMsg("delete user success"))
}
