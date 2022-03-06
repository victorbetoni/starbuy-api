package main

import (
	"authentication-service/database"
	"authentication-service/router"
	"fmt"
	"log"
	"net/http"
)

const port = 5000

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
	fmt.Println("Listening and serving port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}
