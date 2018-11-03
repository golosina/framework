package main

import (
	"github.com/eaperezc/golosina/framework"
	"github.com/eaperezc/golosina/routes"
)

func main() {

	app := framework.New()

	// Prepare routes
	routes.WebRoutes(app.Router)
	routes.APIRoutes(app.Router)

	app.Start()
}
