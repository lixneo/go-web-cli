package routers

import (
	"go-web-cli/controller"
	"go-web-cli/logger"
	"go-web-cli/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 注册zipper日志库
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "GoWeb 后台接口")
	})
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
	}
	return r
}
