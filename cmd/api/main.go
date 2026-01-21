package main

import (
	"log"

	"github.com/lluxy8/todo-app-go/internal/router"
)

func main() {
	r := router.New()

	log.Fatal(r.Run(":8080"))
}
