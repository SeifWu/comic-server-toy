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
	id := c.Param("id")
	iid := c.Param("iid")

	var comicChapter model.ComicChapter
	var previousChapter model.ComicChapter
	var nextChapter model.ComicChapter

	// 上一条
	global.DB.Model(&model.ComicChapter{}).Where("comic_id = ? AND id < ?", id, iid).Order("id desc").Limit(1).Find(&previousChapter)
	// 下一条
	global.DB.Model(&model.ComicChapter{}).Where("comic_id = ? AND id > ?", id, iid).Order("id asc").Limit(1).Find(&nextChapter)

	global.DB.Debug().Find(&comicChapter, iid)

	global.DB.Debug().Model(&comicChapter).Association("ComicChapterDetail").Find(&comicChapter.ComicChapterDetail)

	response.Response(c, http.StatusOK, "0", gin.H{
		"id":           id,
		"iid":          iid,
		"comicChapter": comicChapter,
	}, nil, gin.H{
		// 倒序 换一换
		"previousChapter": nextChapter,
		"nextChapter":     previousChapter,
	})
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
