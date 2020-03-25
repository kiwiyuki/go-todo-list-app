package database

import "github.com/jinzhu/gorm"

type SQLHandler interface {
	Create(value interface{}) *gorm.DB
	Find(out interface{}, where ...interface{}) *gorm.DB
	First(out interface{}, where ...interface{}) *gorm.DB
	Update(after interface{}) *gorm.DB
}
