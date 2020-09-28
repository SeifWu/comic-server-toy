package model

import "gorm.io/gorm"

// Comic 漫画表
type Comic struct {
	gorm.Model

	Name      string `json:"name" gorm:"comment:'漫画名称'"`
	Introduce string `json:"introduce"  gorm:"comment:'简介'"`
	Cover     string `json:"cover" gorm:"comment:'封面'"`
	Author    string `json:"author" gorm:"comment:'作者'"`

	ComicChapter []ComicChapter
}
