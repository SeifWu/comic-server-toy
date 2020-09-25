package v1serializer

import "time"

// UserSerializer 用户管理序列化
type UserSerializer struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Username    string    `json:"username"`
	Nickname    string    `json:"nickname"`
	Avatar      string    `json:"avatar"`
	Mobile      string    `json:"mobile"`
	Description string    `json:"description"`
}
