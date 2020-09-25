package param

// LoginParams 登录参数
type LoginParams struct {
	LoginType string `json:"loginType"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
}
