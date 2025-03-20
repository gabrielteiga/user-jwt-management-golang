package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type IDB interface {
	Open() error
	Close()
	GetDB() *sql.DB
	GetError() error
}

type MySQLDB struct {
	DB    *sql.DB
	Error error
}

func (mysql *MySQLDB) Open() error {
	dsn := fmt.Sprintf(`%s:%s@tcp(127.0.0.1:%s)/%s`, "root", "toor", "3306", "app_db")
	mysql.DB, mysql.Error = sql.Open("mysql", dsn)
	if mysql.Error != nil {
		return mysql.Error
	}

	mysql.Error = mysql.DB.Ping()
	if mysql.Error != nil {
		return mysql.Error
	}

	return nil
}

func (mysql *MySQLDB) Close() {
	mysql.DB.Close()
}

func (mysql *MySQLDB) GetDB() *sql.DB {
	return mysql.DB
}

func (mysql *MySQLDB) GetError() error {
	return mysql.Error
}
