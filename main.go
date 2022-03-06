package main

import (
	"authentication-service/database"
	"authentication-service/router"
	"authentication-service/security"
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

	erroFoda := security.ComparePassword("$2a$10$wmNoJX3C9tyC4Wie9KeIj.v1P6waCEa1omgRBQNMPGFOOJESTqk/i", "lool")
	if erroFoda != nil {
		fmt.Println(erroFoda.Error())
	}

	router := router.Build()
	fmt.Println("Listening and serving port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}
