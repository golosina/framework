package models

import "github.com/golosina/framework/framework"

type Application struct {
	framework.Model
	Name string `json:"name"`
}
