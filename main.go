package main

import (
	"fmt"
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

	fmt.Println("Listening and serving port ", port)

	router := router.Build()
	router.Run(fmt.Sprintf(":%s", port))
}
