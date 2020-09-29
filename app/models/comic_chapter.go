package model

import "gorm.io/gorm"

// ComicChapter 漫画章节
type ComicChapter struct {
	gorm.Model

	Name               string `json:"num" gorm:"comment:'章节'"`
	URL                string `json:"url"  gorm:"comment:'地址'"`
	ComicID            uint   `json:"comic_id"`
	ComicChapterDetail []ComicChapterDetail
}
