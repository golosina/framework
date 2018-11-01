package models

import "github.com/eaperezc/golosina/framework"

type Application struct {
	framework.Model
	Name string `json:"name"`
}
