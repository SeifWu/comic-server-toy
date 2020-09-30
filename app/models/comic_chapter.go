package model

import (
	"database/sql"
	"time"
)

// ComicChapter 漫画章节
type ComicChapter struct {
	ID                 uint                 `json:"id" gorm:"primaryKey"`
	CreatedAt          time.Time            `json:"createdAT"`
	UpdatedAt          time.Time            `json:"updatedAt"`
	DeletedAt          sql.NullTime         `gorm:"index"`
	Name               string               `json:"num" gorm:"comment:'章节'"`
	URL                string               `json:"url"  gorm:"comment:'地址'"`
	ComicID            uint                 `json:"comic_id"`
	ComicChapterDetail []ComicChapterDetail `json:"comicChapterDetail"`
}
