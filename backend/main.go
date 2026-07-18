package main

import (
	"noahsTesting/backend/src/database"
)

func main() {
	//s := routes.NewServer()
	//
	//log.Fatal(s.StartServer(":8080"))

	d := database.New()
	defer d.UsersDatabase.Close()

	d.Login("Noah")

}
