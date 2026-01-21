package main

import (
	"log"
	"todo-app/internal/router"
)

func main() {
	r := router.New()

	log.Fatal(r.Run(":8080"))
}