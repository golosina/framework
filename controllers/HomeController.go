package controllers

import (
	"github.com/golosina/framework/framework"
)

// HomeController definition for the controller
type HomeController struct{}

// Index will render a homepage view
func (c *HomeController) Index(ctx *framework.Context) {

	// Prepare the data we will send to the view, you can define this a a
	// type for HomeViewData in this same file if you want
	viewData := struct {
		Title string
	}{
		Title: "Golosina Home",
	}

	// Render the home view template
	ctx.Response.Render("views/home.html", viewData)
}
