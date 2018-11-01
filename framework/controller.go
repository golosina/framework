package framework

type IResourceController interface {
	Index(*Context)
	Show(*Context)
	Create(*Context)
	Update(*Context)
	Delete(*Context)
}
