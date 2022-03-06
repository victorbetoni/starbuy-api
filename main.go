package main

import (
	"authentication-service/database"
	"authentication-service/router"
	"authentication-service/util"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	util.LoadConfig(".")

	port := os.Getenv("PORT")
	if port == "" {
		port = string(util.GrabConfig().PortAPI)
	}

	err := database.Connect()
	var db = database.GrabDB()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database connection stablished")
	}

	defer db.Close()

	port, err := os.Getenv("PORT")
	if err != nil {
	}

	router := router.Build()
	fmt.Println("Listening and serving port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

}
