package main

import (
	"authentication-service/database"
	"authentication-service/router"
	"fmt"
	"log"
	"net/http"
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

	router := router.Build()
	log.Fatal(http.ListenAndServe(":5000", router))

}
