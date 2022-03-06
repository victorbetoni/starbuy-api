package main

import (
	"authentication-service/router"
	"fmt"
	"log"
	"net/http"
)

const port = 5000

func main() {
	router := router.Build()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
	fmt.Println("Listening and serving port ", port)
}
