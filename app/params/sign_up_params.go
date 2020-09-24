package param

// SignUpParam 注册参数
type SignUpParam struct {
	Username     string `json:"username" label:"用户名"`
	Password     string `json:"password" binding:"required" label:"密码"`
	RePassword   string `json:"rePassword" binding:"required,eqfield=Password" label:"确认密码"`
	Email        string `json:"email" label:"邮箱"`
	RegisterType string `json:"registerType" label:"注册方式"`
	AuthCode     string `json:"authCode" binding:"required" label:"验证码"`
}
