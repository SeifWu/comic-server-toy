package v1manager

import (
	"fmt"
	"net/http"
	"seifwu/app/crawl/qiman6"
	model "seifwu/app/models"
	global "seifwu/global"
	"seifwu/global/response"
	util "seifwu/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var saveParams struct {
	ID string `json:"id" binding:"required"`
}
var qiman qiman6.New

// CrawlComicsController 爬取接口
func CrawlComicsController(c *gin.Context) {
	title := c.Query("title")
	var qiman qiman6.New
	a, _ := qiman.Search(title)

	response.Response(c, http.StatusOK, "0", a.Data, nil, nil)
}

// SaveCrawlComicController 保存爬虫结果
func SaveCrawlComicController(c *gin.Context) {
	// 参数错误处理
	if err := c.ShouldBind(&saveParams); err != nil {
		errResult := util.UnifiedValidation(err, saveParams)

		response.Response(c, http.StatusBadRequest, "40001", nil, errResult, nil)
		return
	}

	comic, _ := qiman.ComicPage(saveParams.ID)

	go saveQiman6Content(comic)

	message := fmt.Sprintf("[%s] 后台努力保存中，详情见漫画管理页面", "comic.Title")
	response.Response(c, http.StatusOK, "0", nil, message, nil)
}

func saveQiman6Content(comic qiman6.Comic) {
	var comicChapters = []model.ComicChapter{}

	for _, value := range comic.ChapterList {
		var comicChapter = model.ComicChapter{
			Name: value.Title,
			URL:  value.URL,
		}

		comicChapters = append(comicChapters, comicChapter)
	}

	var willSave = model.Comic{
		Name:         comic.Title,
		Introduce:    comic.Introduction,
		Cover:        comic.Cover,
		Author:       comic.Author,
		ComicChapter: comicChapters,
	}

	global.DB.FirstOrCreate(&willSave, model.Comic{Name: comic.Title})

	var results = []model.ComicChapter{}

	var count int64
	DB := global.DB.Model(&model.ComicChapter{}).Where("comic_id = ?", 1)

	DB.Count(&count)

	DB.FindInBatches(&results, int(count), func(tx *gorm.DB, batch int) error {
		for _, result := range results {
			pictures, _ := qiman.ComicDetailPage(result.URL)
			var details []model.ComicChapterDetail
			for _, value := range pictures {
				detail := model.ComicChapterDetail{
					Postion: value.Postion,
					URL:     value.URL,
				}
				details = append(details, detail)
			}

			DB := global.DB.Model(&result).Association("ComicChapterDetail")
			var associationCount = DB.Count()
			if int(associationCount) != len(details) {
				global.DB.Model(&result).Association("ComicChapterDetail").Clear()
				global.DB.Model(&result).Association("ComicChapterDetail").Append(details)
				global.DB.Unscoped().Delete(model.ComicChapterDetail{}, "comic_chapter_id IS NULL")
				global.DB.Exec("UPDATE `comic_chapter_details` SET url = REPLACE(url, '//', '') WHERE url LIKE '//%'")
			}
		}

		time.Sleep(5 * time.Second)

		tx.Save(&results)
		return nil
	})
}
