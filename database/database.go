package database

import (
	"authentication-service/util"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Login struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

var db sqlx.DB

// Connect vai criar uma conexão com o banco utilizando as variáveis de ambiente definidas na config
func Connect() (err error) {
	var DBConfig util.Config

	fmt.Println("Starting authentication service...")

	if err = util.LoadConfig(".", &DBConfig); err != nil {
		return
	}

	fmt.Println("Database config loaded from config file")

	//Mantenha o SSLMode ativado, caso contrario ele ficara direcionando para a localhost
	dataSource := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require", DBConfig.HostAddress, DBConfig.Port, DBConfig.Username, DBConfig.Password, DBConfig.Schema)

	var database *sqlx.DB
	if database, err = sqlx.Open(DBConfig.Driver, dataSource); err != nil {
		return err
	}
	db = *database
	return
}

func GrabDB() *sqlx.DB {
	return &db
}
