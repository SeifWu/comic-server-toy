package initializers

import (
	model "seifwu/app/models"
	"seifwu/global"
)

// DBMigrate Migrate Table
func DBMigrate() {
	db := global.DB

	db.AutoMigrate(
		&model.User{},
		&model.Comic{},
		&model.ComicChapter{},
		&model.ComicChapterDetail{},
	)
}
