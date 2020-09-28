package v1manager

import (
	"net/http"
	"seifwu/app/crawl/qiman6"
	"seifwu/global/response"
	util "seifwu/utils"

	"github.com/gin-gonic/gin"
)

var saveParams struct {
	URL string `json:"url" binding:"required" label:"URL地址"`
}

// CrawlComicsController 爬取接口
func CrawlComicsController(c *gin.Context) {
	title := c.Query("title")
	var qiman qiman6.New
	a, _ := qiman.Search(title)

	response.Response(c, http.StatusOK, "0", a.Data, nil, nil)
}

// SaveCrawComicController 保存爬虫结果
func SaveCrawComicController(c *gin.Context) {
	// 参数错误处理
	if err := c.ShouldBind(&saveParams); err != nil {
		errResult := util.UnifiedValidation(err, loginParams)

		response.Response(c, http.StatusBadRequest, "40001", nil, errResult, nil)
		return
	}

	var qiman qiman6.New
	qiman.ComicPage(saveParams.URL)

	response.Response(c, http.StatusOK, "0", "1", nil, nil)
}
