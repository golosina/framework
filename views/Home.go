package views

import (
	"errors"
	"html/template"
	"os"
	"path"

	"github.com/eaperezc/golosina/framework"
)

type HomeView struct {
	Layout   string
	Template string
	Data     interface{}
}

type HomeViewData struct {
	Hello string
	World string
}

func (v *HomeView) Init(ctx *framework.Context) {

	v.Layout = "app"
	v.Template = "home"

	v.Data = &HomeViewData{
		"Good",
		"Night",
	}
}

func (v *HomeView) Render() (*template.Template, error) {

	// prepare layout and template paths
	lp := path.Join("templates/layout", v.Layout+".html")
	fp := path.Join("templates", v.Template+".html")

	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("Template path issue")
		}
	}

	// Log if the file is a directory
	if info.IsDir() {
		return nil, errors.New("Template path is a directory")
	}

	// Parse the templates
	templates, err := template.ParseFiles(lp, fp)
	return templates, err

}

func (v *HomeView) GetLayout() string {
	return v.Layout
}

func (v *HomeView) GetData() interface{} {
	return v.Data
}
