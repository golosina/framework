package framework

import (
	"encoding/json"
	"net/http"
)

// Response is a wrapper for the http.ResponseWriter
type Response struct {
	http.ResponseWriter
}

// String will print the passed string to the response
func (res *Response) String(s string) {
	res.Write([]byte(s))
}

// JSON will try to respond to the client a JSON Object
// we can use. Most common response in APIs for web
func (res *Response) JSON(el interface{}) {
	js, err := json.Marshal(el)
	if err != nil {
		res.String("Boooo")
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}
