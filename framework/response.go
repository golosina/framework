package framework

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	http.ResponseWriter
}

func (res *Response) String(s string) {
	res.Write([]byte(s))
}

func (res *Response) JSON(el interface{}) {
	js, err := json.Marshal(el)
	if err != nil {
		res.String("Boooo")
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}
