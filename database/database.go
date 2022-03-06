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

type User struct {
	Username     string `db:"username"`
	Email        string `db:"email"`
	Name         string `db:"name"`
	Gender       int    `db:"gender"`
	Registration string `db:"registration"`
	Birthdate    string `db:"birthdate"`
	Seller       bool   `db:"seller"`
}

type IncomingUser struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	Gender       int    `json:"gender"`
	Registration string `json:"registration"`
	Birthdate    string `json:"birthdate"`
	Seller       bool   `json:"seller"`
	Password     string `json:"password"`
}

var db sqlx.DB

// Connect vai criar uma conexão com o banco utilizando as variáveis de ambiente definidas na config
func Connect() (err error) {

	config := util.GrabConfig()

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
