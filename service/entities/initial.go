package entities

import (
	"database/sql"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

const (
	dbPath string = "./database.db"
)

var db *sql.db

type SQLExecer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type DaoSource struct {
	SQLExecer
}

func init() {
	isExitst := true
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		isExitst = false
		os.OpenFile(dbPath, os.O_CREATE, os.ModePerm)
	}

	var err error
	db, err = sql.Open("sqlite3", dbPath)
	checkErr(err)
	defer db.Close() //defer

	if !isExitst {
		initTables()
	}

}

func initTables() {
	sqlStmt := `
	create table users(
		id integer primary key auto_increment, 
		username text,
		password text,
		email text,
		phone text
		);
	create table meetings(
		id integer primary key auto_increment,
		title text,
		sponsor text,
		participators text,
		start text,
		end text
		);
	`
	_, err := db.Exec(sqlStmt)
	checkErr(err)

}

func checkErr(err error) {
	if err != nil {
		logs.ErrLog(err)
		panic(err)
	}
}