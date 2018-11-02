package framework

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Request is a wrapper for the http.Request
type Request struct {
	*http.Request
}

// Param is a helper method that can get a parameter
// that is on the url, body or query params for the
// request
func (req *Request) Param(key string) string {
	value := req.QueryParam(key) // first try with the query
	if value == "" {
		value = req.URLParam(key) // try with the url params
	}
	if value == "" {
		value = req.FormParam(key) // form parameters
	}
	if value == "" {
		value = req.JSONFormParam(key) // Json body params
	}
	return value
}

// QueryParam will return the value of a query parameter
func (req *Request) QueryParam(key string) string {
	return req.URL.Query().Get(key)
}

// URLParam returns a parameter that is coming from the
// router request path parameters like /users/{id}
func (req *Request) URLParam(key string) string {
	vars := mux.Vars(req.Request)
	return vars[key]
}

// FormParam returns the query parameters for the body
// of a form post request
func (req *Request) FormParam(key string) string {
	req.ParseForm()        // Parses the request body
	x := req.Form.Get(key) // x will be "" if parameter is not set
	return x
}

// JSONFormParam will get the value for a parameter that
// is part of the body of a request that is sending a
// JSON body and the content type is application/jso
func (req *Request) JSONFormParam(key string) string {
	var f map[string]interface{}
	json.NewDecoder(req.Body).Decode(&f)
	if f[key] != nil {
		return f[key].(string)
	}
	return ""
}

// Validate will help us making sure the parameters that we are
// sending in the requests are valid and for our controller
func (req *Request) Validate(rules map[string]string) (map[string]interface{}, bool) {
	var valid = true
	p := make(map[string]interface{})

	for column, rule := range rules {
		p[column] = req.Param(column)

		// Here we evaluate the available rules
		// that we can use on our parameters
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
