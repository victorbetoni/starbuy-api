package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
	"time"
)

var db sqlx.DB

func Connect() (err error) {
	var database *sqlx.DB
	if database, err = sqlx.Open("postgres", os.Getenv("DATABASE_URL")); err != nil {
		return err
	}
	db = *database
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(time.Minute * 5)
	return
}

func GrabDB() *sqlx.DB {
	if err := db.Ping(); err != nil {
		Connect()
	}
	return &db
}
