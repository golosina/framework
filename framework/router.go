package framework

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct {
	*mux.Router
}

type Context struct {
	Request  *Request
	Response *Response
}

func NewRouter() *Router {
	return &Router{&mux.Router{}}
}

func (r *Router) Get(path string, h func(*Context)) {
	r.makeRequest("GET", path, h)
}

func (r *Router) Post(path string, h func(*Context)) {
	r.makeRequest("POST", path, h)
}

func (r *Router) Put(path string, h func(*Context)) {
	r.makeRequest("PUT", path, h)
}

func (r *Router) Delete(path string, h func(*Context)) {
	r.makeRequest("DELETE", path, h)
}

func (r *Router) Resource(path string, c IResourceController) {
	r.Get(path, c.Index)
	r.Post(path, c.Create)
	r.Put(path, c.Update)
	r.Delete(path, c.Delete)
}

func (r *Router) Group(prefix string, closure func(r *Router)) {
	prefixed := r.PathPrefix(prefix).Subrouter()
	router := &Router{prefixed}
	closure(router)
}

func (r *Router) makeRequest(
	method string,
	path string,
	handler func(*Context),
) {
	r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		context := &Context{&Request{r}, &Response{w}}
		handler(context)
	}).Methods(method)
}
