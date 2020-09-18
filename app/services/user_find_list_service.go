package service

import (
	"fmt"
	"strconv"

	model "seifwu/app/models"
	v1serializer "seifwu/app/serializers/api/v1/manager"
	"seifwu/global"
	util "seifwu/utils"

	"github.com/gin-gonic/gin"
)

// UserFindListService 用户列表服务
func UserFindListService(c *gin.Context, paging bool) gin.H {
	var result []v1serializer.UserSerializer
	var users []model.User
	DB := global.DB.Model(&model.User{})

	if userName, isExist := c.GetQuery("userName"); isExist == true {
		DB = util.LikeCondition(DB, "userName", userName)
	}

	if nickName, isExist := c.GetQuery("nickName"); isExist == true {
		DB = DB.Where("nick_name LIKE ?", fmt.Sprintf("%s%%", nickName))
	}

	total := 0
	DB.Count(&total)

	current, _ := strconv.Atoi(c.Query("current"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	if paging {
		if current == 0 && pageSize == 0 {
			current = 1
			pageSize = 10
		}

		DB = DB.Limit(pageSize).Offset((current - 1) * pageSize)
	}

	DB = DB.Find(&users).Scan(&result)

	return gin.H{
		"success": true,
		"data":    result,
		"pagination": gin.H{
			"current":  current,
			"pageSize": pageSize,
			"total":    total,
		},
	}
}
