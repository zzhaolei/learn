package user

import (
	"simple-user-web/config"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.RouterGroup, config config.Config) {
	user := New(config)
	// 用户
	users := router.Group("/users")
	{
		users.POST("/", user.Create)
		users.GET("/:id", user.Get)
		users.PUT("/:id", user.Update)
		users.DELETE("/:id", user.Delete)
	}
}
