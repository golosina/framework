package controllers

import (
	"log"
	"os"

	"github.com/eaperezc/golosina/framework"
)

type UserController struct {
}

type UserResponse struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func (uc *UserController) Index(c *framework.Context) {
	res := &UserResponse{"Enrique", "32"}
	c.Response.Json(res)
}

func (uc *UserController) Create(c *framework.Context) {
	c.Response.String("user create")
}

func (uc *UserController) Update(c *framework.Context) {
	c.Response.String("user update")
}

func (uc *UserController) Delete(c *framework.Context) {
	log.Println(os.Getenv("APP_NAME"))
	c.Response.String("user delete")
}
