package response

/* Package response 请求规范
{
	"success": true,
	"data": {},
	"errorCode": "0",
	"message": "",
	"meta": {
		"current": 1,
		"pageSize": 10,
		"total": 100,
	}
}
*/

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 响应
func Response(ctx *gin.Context, httpStatus int, errcode string, data interface{}, msg interface{}, meta gin.H) {
	ctx.JSON(
		httpStatus,
		gin.H{
			"success": errcode == "0",
			"errCode": errcode,
			"data":    data,
			"message": msg,
			"meta":    meta,
		},
	)
}

// Success 请求成功
func Success(ctx *gin.Context, data interface{}, msg interface{}, meta gin.H) {
	Response(ctx, http.StatusOK, "0", data, msg, meta)
}

// Fail 请求失败
func Fail(ctx *gin.Context, errcode string, msg interface{}) {
	Response(ctx, http.StatusBadRequest, errcode, nil, msg, nil)
}
