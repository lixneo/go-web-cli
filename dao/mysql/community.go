package mysql

import (
	"database/sql"
	"go-web-cli/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	err = db.Select(&communityList, sqlStr)
	if err == sql.ErrNoRows {
		err = nil
		return
	}
	return
}
