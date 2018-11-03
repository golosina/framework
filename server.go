package main

import (
	"github.com/eaperezc/golosina/controllers"
	"github.com/eaperezc/golosina/framework"
)

func main() {

	app := framework.New()

	app.Router.Resource("/applications", &controllers.ApplicationController{})

	c := &controllers.ApplicationController{}
	app.Router.Get("/test", c.Test)

	app.Router.Group("/api", func(r *framework.Router) {
		c := &controllers.ApplicationController{}
		r.Get("/applications", c.Index)
	})

	app.Start()
}
