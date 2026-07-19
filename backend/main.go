package main

import (
	"log"
	"noahsTesting/backend/src/routes"
)

func main() {
	s := routes.NewServer()

	log.Fatal(s.StartServer(":8080"))

	//d := database.New()
	//defer d.Close()
	//
	//d.Login("Noah")
	//d.DeleteAll()

}
