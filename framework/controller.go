package framework

import (
	"net/http"
)

type IResourceController interface {
	Index(*Context)
	Create(*Context)
	Update(*Context)
	Delete(*Context)
}

type Request struct {
	*http.Request
}

func (req *Request) Param(key string) string {
	return req.URL.Query().Get(key)
}
