package param

// LoginParams 登录参数
type LoginParams struct {
	LoginType string `json:"loginType"`
	UserName  string `json:"userName"`
	PassWord  string `json:"passWord"`
	Email     string `json:"email"`
}
