package database

import (
	"authentication-service/util"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db sqlx.DB

// Connect vai criar uma conexão com o banco utilizando as variáveis de ambiente definidas na config
func Connect() (err error) {

	config := util.GrabConfig()
	fmt.Println(config)

	//Mantenha o SSLMode ativado, caso contrario ele ficara direcionando para a localhost
	dataSource := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require", config.HostAddress, config.Port, config.Username, config.Password, config.Schema)

	var database *sqlx.DB
	if database, err = sqlx.Open(config.Driver, dataSource); err != nil {
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
