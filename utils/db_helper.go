package util

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// LikeCondition Like 查询语句
func LikeCondition(DB *gorm.DB, field string, value interface{}) *gorm.DB {
	formatField := UnderlineString(field)
	likeConditionString := fmt.Sprintf("%s LIKE ?", formatField)
	formatvalue := fmt.Sprintf("%s%%", value)
	return DB.Where(likeConditionString, formatvalue)
}

// EqualCondition 条件语句
func EqualCondition(DB *gorm.DB, field string, value interface{}, condition string) *gorm.DB {
	formatField := UnderlineString(field)
	equalConditionString := fmt.Sprintf("%s %s ?", formatField, condition)
	return DB.Where(equalConditionString, value)
}
