package router

import (
	v1manager "seifwu/app/controllers/api/v1/manager"

	"github.com/gin-gonic/gin"
)

// V1Manager v1 manager 接口
func V1Manager(router *gin.RouterGroup) {
	router.GET("/currentUser", v1manager.CurrentUserController)
	router.GET("/comics", v1manager.FindComicListController)
	router.GET("/comic_crawler", v1manager.CrawlComicsController)
}
