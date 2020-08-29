package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应
func Response(ctx *gin.Context, httpStatus int, code int, data gin.H, msg interface{}) {
	ctx.JSON(
		httpStatus,
		gin.H{
			"code": code,
			"data": data,
			"msg":  msg,
		},
	)
}

// Success 请求成功
func Success(ctx *gin.Context, data gin.H, msg interface{}) {
	Response(ctx, http.StatusOK, 200, data, msg)
}

/*
Fail 请求失败
data 错误返回示例
{
	"success": false,
	"errCode": "40001"
	"errMsg": "错误提示"
}
*/
func Fail(ctx *gin.Context, data gin.H) {
	ctx.JSON(
		http.StatusBadRequest,
		data,
	)
}
