package main

import (
	"log"

	"noahsTesting/pkg/siteTesting"
)

func main() {
	log.Fatal(siteTesting.StartServer(":8080"))
}
