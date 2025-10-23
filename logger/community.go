package logger

import (
	"go-web-cli/dao/mysql"
	"go-web-cli/models"
)

func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}
