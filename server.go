package main

import (
	"github.com/eaperezc/golosina/controllers"
	"github.com/eaperezc/golosina/framework"
)

func main() {

	app := framework.New()

	//c := &controllers.UserController{}
	// app.Router.Get("/users", c.Index)
	// app.Router.Post("/users", c.Create)
	// app.Router.Put("/users", c.Update)
	// app.Router.Delete("/users", c.Delete)

	app.Router.Resource("/users", &controllers.UserController{})

	app.Router.Group("/api", func(r *framework.Router) {
		c := &controllers.UserController{}
		r.Get("/users", c.Index)
	})

	app.Start()
}
