package controller

import (
	"go-web-cli/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	communityList, err := logger.GetCommunityList()
	if err != nil {
		zap.L().Error("mysql.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)
}
