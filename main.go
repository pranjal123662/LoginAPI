package main

import (
	"LoginAPI/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Login API")
	router := router.Router()
	log.Fatal(http.ListenAndServe(":4000", router))
}
