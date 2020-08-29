package param

// SendMailParams 发送邮件参数
type SendMailParams struct {
	Event string `json:"event"`
	Email string `json:"email"`
}
