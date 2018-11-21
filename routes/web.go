package routes

import (
	"github.com/eaperezc/golosina/controllers"
	"github.com/eaperezc/golosina/framework"
)

// WebRoutes will contain the list of all routes
// that we use to show website views
func WebRoutes(r *framework.Router) {

	c := &controllers.HomeController{}
	r.Get("/", c.Index)
}
