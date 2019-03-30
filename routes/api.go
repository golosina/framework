package routes

import (
	"github.com/golosina/framework/controllers"
	"github.com/golosina/framework/framework"
)

// APIRoutes will contain the list of all routes
// that are part of the API
func APIRoutes(r *framework.Router) {

	// We want to prefix the API routes
	r.Group("/api/v1", func(r *framework.Router) {
		c := &controllers.ApplicationController{}
		r.Get("/applications", c.Index)
	})
}
