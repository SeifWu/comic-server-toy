package model

import "github.com/jinzhu/gorm"

// User 用户表
/**
 * TODO
 * UserName UUID 随机
**/
type User struct {
	gorm.Model

	Username    string `json:"username" gorm:"type:varchar(20);unique_index;not null; comment:'用户登录名'"`
	Password    string `json:"-"  gorm:"comment:'用户登录密码'"`
	Email       string `json:"email" gorm:"comment:'邮箱'"`
	NickName    string `json:"nickName" gorm:"comment:'用户昵称'"`
	Avatar      string `json:"avatar" gorm:"comment:'用户头像'"`
	Mobile      string `json:"mobile" gorm:"comment:'联系电话'"`
	Description string `json:"description" gorm:"type:text;comment:'个人描述'"`
	Status      int    `json:"status" gorm:"comment:'0 => 待激活, 1 => 正常'"`
}
