package infra

import (
	"go-todo-list-app/domain/model"
	"go-todo-list-app/interface/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SQLHandler struct {
	Conn *gorm.DB
}

func NewSQLHandler() database.SQLHandler {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err.Error)
	}

	db.AutoMigrate(&model.Task{})

	handler := new(SQLHandler)
	handler.Conn = db
	return handler
}

func (sqlHandler *SQLHandler) Create(value interface{}) *gorm.DB {
	return sqlHandler.Conn.Create(value)
}

func (sqlHandler *SQLHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return sqlHandler.Conn.Find(out, where...)
}

func (sqlHandler *SQLHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return sqlHandler.Conn.First(out, where...)
}
