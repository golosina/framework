# Views

It's really simple to render a view using the framework. For views we are not wrapping any oher library then the core Go templates. If you go to the Home controller you'll see an example on how to render a view from there but here are the basics:

### Controller Code

On our context response struct we have access to a `Render` function that does exacly that, render a template file into the response.

```go
func (c *HomeController) Index(ctx *framework.Context) {
	ctx.Response.Render("views/home.html", nil)
}
```

That should render your home.html view on your browser. Remember to add the route to this handler.

### Passing Variables

The second parameter for the render function is the struct parameters for this particular view. In this case I will show you how to define parameters inside the function.


```go
func (c *HomeController) Index(ctx *framework.Context) {

	// Prepare the data we will send to the view, you can define this a a
	// type for HomeViewData in this same file if you want
	viewData := struct {
		Title string
	}{
		Title: "Golosina Home",
	}

	// Render the home view template
	ctx.Response.Render("views/home.html", viewData)
}
```

Of course you could also create a `type` for the viewData like this:

```go
type HomeViewData struct {
	Title string
}

// Index will render a homepage view
func (c *HomeController) Index(ctx *framework.Context) {

	viewData := &HomeViewData{
		Title: "Golosina Home",
	}

	// Render the home view template
	ctx.Response.Render("views/home.html", viewData)
}
```

With this (any of the examples above) you will have access to the struct public properties and you can use them like this:

```html
<title>{{.Title}}</title>
```

?> We use a `views` folder to store all the templates we want to render in our app. You can potentially use any folder you want but we recommend to keep them there for consistency.
