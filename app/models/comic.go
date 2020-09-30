package model

import (
	"database/sql"
	"time"
)

// Comic 漫画表
type Comic struct {
	ID        uint         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time    `json:"createdAT"`
	UpdatedAt time.Time    `json:"updatedAt"`
	DeletedAt sql.NullTime `gorm:"index"`
	Name      string       `json:"name" gorm:"comment:'漫画名称'"`
	Introduce string       `json:"introduce"  gorm:"comment:'简介'"`
	Cover     string       `json:"cover" gorm:"comment:'封面'"`
	Author    string       `json:"author" gorm:"comment:'作者'"`

	ComicChapter []ComicChapter `json:"comicChapter"`
}
