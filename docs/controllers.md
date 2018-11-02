
# Controllers

As we are all used to we have controller files where we will run the operations that we listed in our routes. Defining a controller in `Golosina` is simple. Lets take a look:

## Definition

First we need to create a file in out `controllers` folder that will contain the definition of our controller struct:

```go
package controllers

type UserController struct {}
```

## Responses

Since we need to prepare the response we are going to send back to the cliend we need to define a struct for that, here's an example of an API response for the `User` model:

```go
type UserResponse struct {
	Success     bool            `json:"success"`
	User        *models.User    `json:"user,omitempty"`
	Users       []*models.User  `json:"users,omitempty"`
	Message     string          `json:"message,omitempty"`
}
```

This is not set on stone. You can define the responses however you want to.

?> You can define as many responses structs you need per controller, we just propose that you keep them inside their own controller file so it's easier to maintain later.

## Actions

Controller functions will work as our handlers for the request operations. Here's a simple example to return "OK":

```go
func (c *ExampleController) Health(ctx *framework.Context) {
    ctx.Response.String("OK")
}
```

Simple right? Now lets see an more complex example just to get the idea of what we can do with the responses we defined earlier. Here we will see how we can use Models and JSON responses the way we normally would want to:

```go
func (c *UserController) Index(ctx *framework.Context) {

    var users []*models.User
	ctx.Database.Find(&users)

    // Return JSON to the client
	ctx.Response.JSON(&UserResponse{
		Success:    true,
		Users:      Users,
	})
}
```

?> Read the `Context` documentation to see what we have available on our controller action handlers.