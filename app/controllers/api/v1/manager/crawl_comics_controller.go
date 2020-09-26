package v1manager

import (
	"net/http"
	"seifwu/app/crawl/qiman6"
	"seifwu/global/response"

	"github.com/gin-gonic/gin"
)

// CrawlComicsController 爬取接口
func CrawlComicsController(c *gin.Context) {
	var qiman qiman6.New
	a, _ := qiman.Search("鬼灭之刃")

	response.Response(c, http.StatusOK, "0", a.Data, nil, nil)
}
