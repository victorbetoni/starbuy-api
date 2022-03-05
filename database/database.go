package database

import (
	"authentication-service/util"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Login struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

func Connect(db *sqlx.DB) (err error) {
	var DBConfig util.Config

	fmt.Println("Starting authentication service...")

	if err = util.LoadConfig(".", &DBConfig); err != nil {
		return
	}

	fmt.Println("Database config loaded from config file")

	//Mantenha o SSLMode ativado, caso contrario ele ficara direcionando para a localhost
	dataSource := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require", DBConfig.HostAddress, DBConfig.Port, DBConfig.Username, DBConfig.Password, DBConfig.Schema)

	database, err := sqlx.Open(DBConfig.Driver, dataSource)
	*db = *database
	if err != nil {
		return
	}
	return
}
