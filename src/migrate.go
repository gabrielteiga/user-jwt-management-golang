package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

var dsn = "root:toor@tcp(127.0.0.1:3306)/app_db"
var migrationDir = "migrations"

func main() {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database -> ", err)
	}
	defer db.Close()

	files, err := os.ReadDir(migrationDir)
	if err != nil {
		log.Fatal("Error reading migrations dir -> ", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			err := executableSQLFile(db, filepath.Join(migrationDir, file.Name()))
			if err != nil {
				log.Printf("Error running %s -> %v", file.Name(), err)
			} else {
				log.Printf("Migration %s finished successfully", file.Name())
			}
		}
	}
}

type DBMHasAlreadyBeenExecuted struct {
	Message string
}

func (dbm *DBMHasAlreadyBeenExecuted) Error() string {
	return "The file has already been executed in the database. No tables and rows affected."
}

type DBMSaveCouldNotBeExecuted struct {
	Message string
}

func (dbm *DBMSaveCouldNotBeExecuted) Error() string {
	return "Something went wrong in the dbm execution."
}

func executableSQLFile(db *sql.DB, fileName string) error {
	if fileHasBeenExecuted(db, fileName) {
		return &DBMHasAlreadyBeenExecuted{}
	}

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(fileContent))
	if err != nil {
		return err
	}

	err = saveToDBM(db, fileName)
	if err != nil {
		return err
	}

	return nil
}

func fileHasBeenExecuted(db *sql.DB, fileName string) bool {
	var fileTmp string
	sql := "SELECT file FROM dbm_control WHERE file = ?"

	err := db.QueryRow(sql, fileName).Scan(&fileTmp)
	if err != nil {
		return false
	}

	return true
}

func saveToDBM(db *sql.DB, fileName string) error {
	sql := "INSERT INTO dbm_control (file) VALUES(?)"

	_, err := db.Exec(sql, fileName)
	if err != nil {
		return &DBMSaveCouldNotBeExecuted{}
	}

	return nil
}
