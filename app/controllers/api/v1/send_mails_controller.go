package v1api

import (
	"time"

	param "seifwu/app/params"
	"seifwu/global"
	"seifwu/global/response"
	util "seifwu/utils"

	"github.com/gin-gonic/gin"
)

func mailContent(authCode string) string {
	// TODO：将来改为字符串模板
	content := "<div style='padding: 10px; text-align: center;'>" +
		"<span style='color: #0050b3; font-size: 22px'>" +
		authCode +
		"</span>" +
		"<p>验证码 2 分钟内有效</p>" +
		"</div>"
	return content
}

// SendAuthCodeMailsController 发送验证码邮件
func SendAuthCodeMailsController(c *gin.Context) {
	var sendMailParams param.SendMailParams
	rdb := global.RDB0

	err := c.ShouldBindJSON(&sendMailParams)
	if err != nil {
		response.Fail(c, "40004", "传递参数有误")
		return
	}

	mailTo := []string{
		sendMailParams.Email,
	}
	authCode := util.RandomAuthCodeString(6)

	if sendMailParams.Event == "register" {
		err = util.SendMail(mailTo, "[我猜你在找验证码]", mailContent(authCode))
		if err != nil {
			response.Fail(c, "40004", err.Error())
			return
		}

		err = rdb.Set(c, sendMailParams.Email+"-register", authCode, 120*time.Second).Err()
		if err != nil {
			response.Fail(c, "40004", err)
			return
		}
	}

	response.Success(c, gin.H{"user": "x", "email": sendMailParams.Email}, "1", nil)
}
