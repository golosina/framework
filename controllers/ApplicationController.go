package controllers

import (
	"github.com/eaperezc/golosina/framework"
	"github.com/eaperezc/golosina/models"
)

type ApplicationController struct {
}

type ApplicationResponse struct {
	Success      bool                  `json:"success"`
	Application  *models.Application   `json:"application,omitempty"`
	Applications []*models.Application `json:"applications,omitempty"`
	Message      string                `json:"message,omitempty"`
}

func (c *ApplicationController) Index(ctx *framework.Context) {
	var applications []*models.Application
	ctx.Database.Debug().Find(&applications)

	ctx.Response.JSON(&ApplicationResponse{
		Success:      true,
		Applications: applications,
	})
}

func (c *ApplicationController) Show(ctx *framework.Context) {

	params, valid := ctx.Request.Validate(map[string]string{
		"id": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	var application models.Application
	ctx.Database.Debug().First(&application, params["id"]) // find app with id 1

	if application.ID == 0 {
		ctx.Response.JSON(&ApplicationResponse{Success: false, Message: "Application doesn't exist"})
		return
	}
	ctx.Response.JSON(&ApplicationResponse{Success: true, Application: &application})
}

func (c *ApplicationController) Create(ctx *framework.Context) {

	params, valid := ctx.Request.Validate(map[string]string{
		"name": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	application := models.Application{
		Name: params["name"].(string),
	}

	ctx.Database.Debug().Create(&application)
	ctx.Response.JSON(&ApplicationResponse{Success: true, Application: &application})
}

func (c *ApplicationController) Update(ctx *framework.Context) {

	params, valid := ctx.Request.Validate(map[string]string{
		"id":   "required",
		"name": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	var application models.Application
	ctx.Database.First(&application, params["id"])
	ctx.Database.Model(&application).Update("name", params["name"])

	ctx.Response.JSON(&ApplicationResponse{Success: true, Application: &application})
}

func (c *ApplicationController) Delete(ctx *framework.Context) {
	params, valid := ctx.Request.Validate(map[string]string{
		"id": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	var application models.Application
	ctx.Database.First(&application, params["id"])
	ctx.Database.Delete(&application)

	ctx.Response.JSON(&ApplicationResponse{Success: true})

}
