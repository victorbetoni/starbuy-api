package main

import (
	"authentication-service/database"
	"fmt"
)

func main() {

	var err = database.Connect()
	var db = database.GrabDB()
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
