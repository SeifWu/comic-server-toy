package scope

import "gorm.io/gorm"

// UserFindByUsername 通过用户名寻找用户
func UserFindByUsername(username string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("username = ?", username)
	}
}
