package main

import (
	"example_middleware/database"
	"example_middleware/router"
)

func main() {
	database.StartDB()
	r := router.StarrtApp()
	r.Run(":8086")
}