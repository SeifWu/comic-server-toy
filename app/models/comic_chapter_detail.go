package model

import "github.com/jinzhu/gorm"

// ComicChapterDetail 漫画内容
type ComicChapterDetail struct {
	gorm.Model

	Postion        int    `gorm:"autoIncrement"`
	URL            string `json:"img"`
	ComicChapterID uint
}
