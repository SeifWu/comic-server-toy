package v1manager

import (
	"net/http"
	"seifwu/app/crawl/qiman6"
	"seifwu/global/response"

	"github.com/gin-gonic/gin"
)

// CrawlComicsController 爬取接口
func CrawlComicsController(c *gin.Context) {
	var qiman qiman6.Qiman6
	// a, _ := qiman.Search("鬼灭之刃")
	// qiman.ComicPageQiman6("http://www.qiman6.com/16384/")
	qiman.ComicChapterListQiman6("http://www.qiman6.com/16384/")
	// fmt.Println("result: ", a)
	response.Response(c, http.StatusOK, "0", "a", nil, nil)
}
