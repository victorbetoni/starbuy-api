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
	Username     string `db:"username" json:"username"`
	Email        string `db:"email" json:"email"`
	Name         string `db:"name" json:"name"`
	Gender       int    `db:"gender" json:"gender"`
	Registration string `db:"registration" json:"registration"`
	Birthdate    string `db:"birthdate" json:"birthdate"`
	Seller       bool   `db:"seller" json:"seller"`
}

type PasswordBindedUser struct {
	User
	Password string
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
	if err := db.Ping(); err != nil {
		Connect()
	}
	return &db
}
