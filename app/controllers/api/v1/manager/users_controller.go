package v1manager

import (
	"net/http"

	service "seifwu/app/services"
	"seifwu/global/response"

	"github.com/gin-gonic/gin"
)

// FindListUser GET /api/v1/manager/user
func FindListUser(c *gin.Context) {
	result := service.UserFindListService(c, true)
	if result["success"] == false {
		response.Fail(c, result)
		return
	}

	c.JSON(
		http.StatusOK,
		result,
	)
}
