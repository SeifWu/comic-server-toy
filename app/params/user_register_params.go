package param

// RegisterParams 注册参数
type RegisterParams struct {
	Username     string `json:"userName"`
	Password     string `json:"passWord"`
	Email        string `json:"email"`
	RegisterType string `json:"register_type"`
	AuthCode     string `json:"authCode"`
}
