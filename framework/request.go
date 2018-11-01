package framework

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Request struct {
	*http.Request
}

func (req *Request) Param(key string) string {
	value := req.QueryParam(key)
	if value == "" {
		value = req.URLParam(key)
	}
	if value == "" {
		value = req.FormParam(key)
	}
	if value == "" {
		value = req.JSONFormParam(key)
	}
	return value
}

func (req *Request) QueryParam(key string) string {
	return req.URL.Query().Get(key)
}

func (req *Request) URLParam(key string) string {
	vars := mux.Vars(req.Request)
	return vars[key]
}

func (req *Request) FormParam(key string) string {
	req.ParseForm()        // Parses the request body
	x := req.Form.Get(key) // x will be "" if parameter is not set
	return x
}

func (req *Request) JSONFormParam(key string) string {
	var f map[string]interface{}
	json.NewDecoder(req.Body).Decode(&f)
	if f[key] != nil {
		return f[key].(string)
	}
	return ""
}

func (req *Request) Validate(rules map[string]string) (map[string]interface{}, bool) {

	var valid = true
	p := make(map[string]interface{})

	for column, rule := range rules {

		p[column] = req.Param(column)

		switch rule {
		case "required":
			if p[column] == "" {
				log.Println(p[column])
				valid = false
			}
		}
	}
	return p, valid
}
