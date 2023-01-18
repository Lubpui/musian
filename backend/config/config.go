package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect_Database() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "root:root1234@tcp(musian-database.cyb2u3y66i8t.ap-southeast-1.rds.amazonaws.com:3306)/musian")
	if err != nil {
		log.Fatal("error:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error:", err)
	}
	return db
}
