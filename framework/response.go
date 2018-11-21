package framework

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
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

// Render will render a template into the Response, it requires the
// string path of the template inside the views folder and the
// parameters we want to sent to it
func (res *Response) Render(fp string, params interface{}) {

	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("Template path issue")
		}
	}

	// Log if the file is a directory
	if info.IsDir() {
		log.Println("Template path is a directory")
	}

	t, err := template.ParseFiles(fp) // Parse template file.
	if err != nil {
		log.Println(err.Error())
	}

	t.Execute(res, params) // merge.
}
