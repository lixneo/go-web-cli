package controller

import (
	"go-web-cli/dao/mysql"
	"go-web-cli/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {
	// 获取请求参数,并校验
	var fo models.ParamSignUp
	if err := c.ShouldBindJSON(&fo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  err.Error(),
		})
		return
	}
	// 业务处理
	err := mysql.Register(&models.User{
		UserName: fo.Username,
		Password: fo.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
	})
}
