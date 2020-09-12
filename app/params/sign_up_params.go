package param

type SignUpParam struct {
	Username     string `json:"userName" `
	Password     string `json:"passWord"`
	Email        string `json:"email"`
	RegisterType string `json:"registerType"`
	AuthCode     string `json:"authCode"`
}
