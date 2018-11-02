# Context

The context is the struct instance that we send to our Controller actions. With this you have access to various elements in the request flow. Take a look at the definition:

```go
type Context struct {
	Request  *Request
	Response *Response
	Database *Database
}
```

## Request

The request is a wrapper for the internal `http.Request` so we have access to all of the functions available for it. Using it we have access to some helper functions to get data from our requests.

#### Query Parameters

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    value := req.QueryParam(key)
    // do something with the query param value
}
```

#### URL Parameters

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    value = req.URLParam(key)
    // do something with the URL param value
}
```

#### JSON Parameters

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    value = req.JSONFormParam(key)
    // do something with the JSON param value
}
```

#### Form Parameters

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    value = req.FormParam(key)
    // do something with the Form param value
}
```

#### Any Parameter

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    value = req.Param(column)
    // do something with the Form param value
}
```

## Response

A Response is a wrapper for the internal `http.ResponseWriter` so we have access to all of the functions available for it. Using it we have access to some helper functions to send data back to the client.

#### Print Strings

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    ctx.Response.String("This will be written to the response")
}
```

#### JSON

```go
type ApplicationResponse struct {
	Success    bool    `json:"success"`
	Data       *Data   `json:"data"`
}

func (c *ExampleController) Index(ctx *framework.Context) {
    var data Data
    ctx.Response.JSON(&ExampleResponse{Success: true, Data: &data})
}
```

This will return a json response like this:

```js
{
    "success": true,
    "data": "<your data>"
}
```

## Database

We initialize the database for you using the configurations you have on the `.env` file. This struct will give you access to all the GORM database functions.

```go
func (c *ExampleController) Index(ctx *framework.Context) {
    var example models.Example
    ctx.Database.Debug().First(&example, 1)
    // do something with the example 1
}
```

?> For more information on how to use the `Database` and available functions please check out the official [GORM Documentation](http://gorm.io/docs/). There are also some examples on how this work on the Models guide.
