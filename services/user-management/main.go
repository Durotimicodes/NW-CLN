package main

import (
	"github.com/durotimicodes/natwest-clone/user-service/routes"
	"github.com/durotimicodes/natwest-clone/user-service/server"
	"gorm.io/gorm"
)

func main() {


	//Initialize the database
	//db := database.InitDB()
	var db *gorm.DB //use JSON file for now //saving to disc

	//Set up routes with the database
	router := routes.SetUpUserRoutes(db)

	//Start the server
	server.SetUpUserServer(router)

}
