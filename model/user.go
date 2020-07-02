package model

import "github.com/jinzhu/gorm"

// User 用户表
type User struct {
	gorm.Model

	UserName  string `json:"userName" gorm:"type:varchar(20); not null; comment:'用户登录名'"`
	Password  string `json:"-"  gorm:"comment:'用户登录密码'"`
	NickName  string `json:"nickName" gorm:"comment:'用户昵称'"`
	HeaderImg string `json:"headerImg" gorm:"default:'https://i2.hdslb.com/bfs/face/dcb38241d6aa5279a54c663b171d81cff067c087.jpg@70w_70h_1c_100q.webp';comment:'用户头像'"`
}
