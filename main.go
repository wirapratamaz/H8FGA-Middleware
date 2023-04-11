package main

import (
	"jwtgo/database"
	"jwtgo/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
