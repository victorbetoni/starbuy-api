package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"starbuy/database"
	"starbuy/router"
	"starbuy/util"

	_ "github.com/lib/pq"
)

func main() {

	util.LoadConfig(".")

	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprint(util.GrabConfig().PortAPI)
	}

	var err = database.Connect()
	var db = database.GrabDB()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database connection stablished")
	}
	err = nil
	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()

	router := router.Build()
	fmt.Println("Listening and serving port ", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))

}
