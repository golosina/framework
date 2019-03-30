package routes

import (
	"github.com/golosina/framework/controllers"
	"github.com/golosina/framework/framework"
)

// WebRoutes will contain the list of all routes
// that we use to show website views
func WebRoutes(r *framework.Router) {

	c := &controllers.HomeController{}
	r.Get("/", c.Index)
}
