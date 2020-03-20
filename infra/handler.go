package infra

import (
	"database/sql"
	"go-todo-list-app/interface/database"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type SQLHandler struct {
	Conn *sql.DB
}

func NewSQLHandler() database.SQLHandler {
	conn, err := sql.Open("sqlite3", "./task.db")
	if err != nil {
		panic(err.Error)
	}
	handler := new(SQLHandler)
	handler.Conn = conn
	return handler
}

func (sqlHandler *SQLHandler) Execute(query string, args ...interface{}) (database.Result, error) {
	res := SQLResult{}
	result, err := sqlHandler.Conn.Exec(query, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (sqlHandler *SQLHandler) Query(query string, args ...interface{}) (database.Row, error) {
	rows, err := sqlHandler.Conn.Query(query, args...)
	if err != nil {
		return new(SQLRow), err
	}
	row := new(SQLRow)
	row.Rows = rows
	return row, nil
}

type SQLResult struct {
	Result sql.Result
}

func (sqlResult SQLResult) LastInsertId() (int64, error) {
	return sqlResult.Result.LastInsertId()
}

func (sqlResult SQLResult) RowsAffected() (int64, error) {
	return sqlResult.Result.RowsAffected()
}

type SQLRow struct {
	Rows *sql.Rows
}

func (sqlRow *SQLRow) Scan(args ...interface{}) error {
	return sqlRow.Rows.Scan(args)
}

func (sqlRow *SQLRow) Next() bool {
	return sqlRow.Rows.Next()
}

func (sqlRow *SQLRow) Close() error {
	return sqlRow.Rows.Close()
}
