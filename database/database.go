package database

import (
	"os"
	"starbuy/util"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db sqlx.DB

func Connect() (err error) {

	config := util.GrabConfig()

	//Mantenha o SSLMode ativado, caso contrario ele ficara direcionando para a localhost
	/*dataSource := fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=require", config.HostAddress, config.Port, config.Username, config.Password, config.Schema)*/

	var database *sqlx.DB
	if database, err = sqlx.Open(config.Driver, os.Getenv("DATABASE_URL")); err != nil {
		return err
	}
	db = *database
	return
}

func GrabDB() *sqlx.DB {
	if err := db.Ping(); err != nil {
		Connect()
	}
	return &db
}
