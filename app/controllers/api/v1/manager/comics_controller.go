package v1manager

import (
	"net/http"
	"seifwu/global/response"

	"github.com/gin-gonic/gin"
)

// FindComicListController 漫画列表
func FindComicListController(c *gin.Context) {
	response.Response(c, http.StatusOK, "0", gin.H{"test": "请求成功"}, nil, nil)
}

// FindComicController 漫画详情
func FindComicController(c *gin.Context) {

}

// CreateComicController 创建漫画
func CreateComicController(c *gin.Context) {

}

// UpdateComicController 更新漫画
func UpdateComicController(c *gin.Context) {

}

// DeleteComicController 删除漫画
func DeleteComicController(c *gin.Context) {

}
