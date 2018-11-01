# Golosina

Simple web MVC framework build on go.

# Initialization

```go
func main() {
	app := framework.New()
	app.Router.Get("/", YourHandler)
	app.Start()
}
```

# Routing

Based on gorilla/mux router.

### By Method Type

```go
    c := &controllers.UsersController{}

    app.Router.Get("/users", c.Index)
    app.Router.Post("/users", c.Create)
    app.Router.Get("/users/{id}", c.Show)
    app.Router.Put("/users/{id}", c.Update)
    app.Router.Delete("/users/{id}", c.Delete)
```

### Resource Routes

```go
    // Or this will do the same as above
	app.Router.Resource("/users", &controllers.UserController{})
```

### Prefixes

```go
	app.Router.Group("/api", func(r *framework.Router) {
		c := &controllers.UserController{}
		r.Get("/rsers", c.Index)
    })
```

# Controllers

This is the structure of a basic Resource controller

### Define responses

```go
type UserController struct {
}

type UserResponse struct {
	Success     bool            `json:"success"`
	User        *models.User    `json:"user,omitempty"`
	Users       []*models.User  `json:"users,omitempty"`
	Message     string          `json:"message,omitempty"`
}
```

### List

```go
func (c *UserController) Index(ctx *framework.Context) {

    var users []*models.User
	ctx.Database.Find(&users)

	ctx.Response.JSON(&UserResponse{
		Success:    true,
		Users:      Users,
	})
}
```

### Find By Id

```go
func (c *UserController) Show(ctx *framework.Context) {

	params, valid := ctx.Request.Validate(map[string]string{
		"id": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	var User models.User
	ctx.Database.First(&user, params["id"])

	if User.ID == 0 {
		ctx.Response.JSON(&UserResponse{Success: false, Message: "User doesn't exist"})
		return
	}
	ctx.Response.JSON(&UserResponse{Success: true, User: &user})
}
```

### Create

```go
func (c *UserController) Create(ctx *framework.Context) {

	params, valid := ctx.Request.Validate(map[string]string{
		"name": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	user := models.User{
		Name: params["name"].(string),
	}

	ctx.Database.Create(&user)
	ctx.Response.JSON(&UserResponse{Success: true, User: &user})
}
```

### Update

```go
func (c *UserController) Update(ctx *framework.Context) {

	params, valid := ctx.Request.Validate(map[string]string{
		"id":   "required",
		"name": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	var User models.User
	ctx.Database.First(&user, params["id"])

	if User.ID == 0 {
		ctx.Response.JSON(&UserResponse{Success: false, Message: "User doesn't exist"})
		return
	}

	ctx.Database.Model(&user).Update("name", params["name"])
	ctx.Response.JSON(&UserResponse{Success: true, User: &user})
}
```

### Delete

```go
func (c *UserController) Delete(ctx *framework.Context) {
	params, valid := ctx.Request.Validate(map[string]string{
		"id": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	var User models.User
	ctx.Database.First(&user, params["id"])

	if User.ID == 0 {
		ctx.Response.JSON(&UserResponse{Success: false, Message: "User doesn't exist"})
		return
	}

	ctx.Database.Delete(&user)

	ctx.Response.JSON(&UserResponse{Success: true})

}
```
