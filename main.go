package main

import (
	"authentication-service/util"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {

	var DBConfig util.Config

	fmt.Println("Starting authentication service...")

	if err := util.LoadConfig(".", &DBConfig); err != nil {
		panic(err)
	}

	fmt.Println("Database config loaded from config file")

	dataSource := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", DBConfig.Host, DBConfig.Port, DBConfig.Username, DBConfig.Password, DBConfig.Schema)

	db, err = sql.Open(DBConfig.Driver, dataSource)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database connection stablished")
	}

	defer db.Close()

	error1 := db.Ping()
	if error1 != nil {
		fmt.Println(error1.Error())
	}

	var lines *sql.Rows
	if lines, err = db.Query("SELECT * FROM products"); err != nil {
		fmt.Println("Entrou no")
		panic(err.Error())
	}

	fmt.Println(lines)

}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
