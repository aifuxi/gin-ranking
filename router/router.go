package router

import (
	"github.com/aifuxi/gin-ranking/controller"
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			userGroup.GET("/batch", controller.UserController{}.Batch)
			userGroup.POST("", controller.UserController{}.Create)
			userGroup.PUT("", controller.UserController{}.Update)
			userGroup.DELETE("", controller.UserController{}.Delete)
		}
	}

	return r
}
