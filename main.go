package main

import (
	"authentication-service/database"
	"authentication-service/util"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func main() {

	var err error
	var DBConfig util.Config

	fmt.Println("Starting authentication service...")

	if err := util.LoadConfig(".", &DBConfig); err != nil {
		panic(err)
	}

	fmt.Println("Database config loaded from config file")

	//Mantenha o SSLMode ativado, caso contrario ele ficara direcionando para a localhost
	dataSource := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=require", DBConfig.HostAddress, DBConfig.Port, DBConfig.Username, DBConfig.Password, DBConfig.Schema)

	db, err = sqlx.Open(DBConfig.Driver, dataSource)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database connection stablished")
	}

	defer db.Close()

	logins := []database.Login{}

	err = db.Select(&logins, "SELECT * FROM login")
	if err != nil {
		panic(err.Error())
	}

	for _, row := range logins {
		fmt.Println(row)
	}

}
