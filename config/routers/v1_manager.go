package router

import (
	v1manager "seifwu/app/controllers/api/v1/manager"

	"github.com/gin-gonic/gin"
)

// V1Manager v1 manager 接口
func V1Manager(router *gin.RouterGroup) {
	router.GET("/comics", v1manager.FindComicListController)
	router.GET("/craw_comic", v1manager.CrawlComicsController)
}
