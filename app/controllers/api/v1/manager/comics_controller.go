package v1manager

import (
	"net/http"
	model "seifwu/app/models"
	"seifwu/global"
	"seifwu/global/response"
	scope "seifwu/global/scopes"

	"github.com/gin-gonic/gin"
)

// FindComicListController 漫画列表
func FindComicListController(c *gin.Context) {
	var comics []model.Comic
	// Table("comics")
	global.DB.Model(&model.Comic{}).Preload("ComicChapter").Scopes(scope.Paginate(c)).Find(&comics)
	response.Response(c, http.StatusOK, "0", comics, nil, nil)
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
