package infra

import (
	"fmt"
	"go-todo-list-app/domain/model"
	"go-todo-list-app/interface/database"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type SQLHandler struct {
	Conn *gorm.DB
}

func NewSQLHandler() database.SQLHandler {
	env, err := godotenv.Read("/app/.env")
	if err != nil {
		panic(err.Error)
	}

	protocol := "tcp(database:3306)"
	params := "charset=utf8mb4&parseTime=True"
	conn := fmt.Sprintf("%s:%s@%s/%s?%s", env["MYSQL_USER"], env["MYSQL_PASSWORD"], protocol, env["MYSQL_DATABASE"], params)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		panic(err.Error)
	}
	db.LogMode(true)
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
