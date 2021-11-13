package main

import (
	"github.com/rohandas-max/admybrand/database"
	"github.com/rohandas-max/admybrand/router"
)

func main() {
	database.Connection()

	routes := router.Router()

	routes.Run(":4000")

}
