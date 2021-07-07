package main

import (
	"ToDo/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	var port string
	fmt.Println("Enter the PORT: ")
	fmt.Scanln(&port)
	fmt.Printf("Starting server on the port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
