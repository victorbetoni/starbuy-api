package main

import (
	"authentication-service/database"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db sqlx.DB

func main() {

	var err = database.Connect(&db)
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
