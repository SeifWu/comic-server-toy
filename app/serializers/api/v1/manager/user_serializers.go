package managerv1serializer

// ResultSerializer 用户管理序列化
type ResultSerializer struct {
	UserName    string `json:"userName"`
	NickName    string `json:"nickName"`
	Avatar      string `json:"avatar"`
	Mobile      string `json:"mobile"`
	Description string `json:"description"`
}
