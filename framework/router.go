package framework

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	mux.Router
}

type Context struct {
	Request  *Request
	Response *Response
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Get(path string, h func(*Context)) {
	r.MakeRequest("GET", path, h)
}

func (r *Router) Post(path string, h func(*Context)) {
	r.MakeRequest("POST", path, h)
}

func (r *Router) Put(path string, h func(*Context)) {
	r.MakeRequest("PUT", path, h)
}

func (r *Router) Delete(path string, h func(*Context)) {
	r.MakeRequest("DELETE", path, h)
}

func (r *Router) Resource(path string, c IController) {
	r.Get(path, c.Index)
	r.Post(path, c.Create)
	r.Put(path, c.Update)
	r.Delete(path, c.Delete)
}

func (r *Router) MakeRequest(
	method string,
	path string,
	handler func(*Context),
) {
	r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		context := &Context{&Request{r}, &Response{w}}
		handler(context)
	}).Methods(method)
}
