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
	fmt.Println(port)
	if port == "" {
		port = fmt.Sprint(util.GrabConfig().PortAPI)
	}
	fmt.Println(os.Getenv("DATABASE_URL"))

	err := database.Connect()
	var db = database.GrabDB()
	db.Ping()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database connection stablished")
	}

	defer db.Close()

	router := router.Build()
	fmt.Println("Listening and serving port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))

}
