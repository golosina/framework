package framework

// IResourceController is the interface for all our
// controllers that are using the Resource routes
type IResourceController interface {
	Index(*Context)
	Show(*Context)
	Create(*Context)
	Update(*Context)
	Delete(*Context)
}
