package v1public

import (
	"net/http"
	"seifwu/global/response"

	"github.com/gin-gonic/gin"
)

// FindComicListController 漫画列表
func FindComicListController(c *gin.Context) {
	response.Response(c, http.StatusOK, "0", gin.H{"test": "请求成功"}, nil, nil)
}
