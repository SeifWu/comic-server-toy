package model

import (
	"database/sql"
	"time"
)

// ComicChapterDetail 漫画内容
type ComicChapterDetail struct {
	ID             uint         `json:"id" gorm:"primaryKey"`
	CreatedAt      time.Time    `json:"createdAT"`
	UpdatedAt      time.Time    `json:"updatedAt"`
	DeletedAt      sql.NullTime `gorm:"index"`
	Postion        int          `gorm:"autoIncrement"`
	URL            string       `json:"img"`
	ComicChapterID uint
}
