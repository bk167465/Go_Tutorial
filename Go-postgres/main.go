package main

import (
	"fmt"
	"github.com/bk167465/Go_Tutorial/Go-postgres/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
