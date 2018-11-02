# Models

The models will help us do the actual logic on our Database entities. The models on the framework work as a wrapper for the [GORM Models](http://gorm.io/docs/), please make sure to check out their documentation.

## Query

To get the data from our tables we have a couple of methods we can use depending on what we want to get. These are just examples on how you can achieve simple Resource actions with the models.

### Collection of Models

Here's an example of an API controller action to get all Users from our database and return it as a JSON response.

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

### Find Model By Id

Our resource routes have an action to get a model by it's ID, heres an example on how we can get the `User` by it's ID.

```go
func (c *UserController) Show(ctx *framework.Context) {

    // Validate request parameters
	params, valid := ctx.Request.Validate(map[string]string{
		"id": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

    // Find our user
	var User models.User
	ctx.Database.First(&user, params["id"])

    // Respond
	if User.ID == 0 {
		ctx.Response.JSON(&UserResponse{Success: false, Message: "User doesn't exist"})
		return
	}
	ctx.Response.JSON(&UserResponse{Success: true, User: &user})
}
```


## Create

When we are creating new users we need some data from the request body, URL parameters, etc. Here's an example on how you can validate and use the parameters from the request:

```go
func (c *UserController) Create(ctx *framework.Context) {

    // Validate request parameters
	params, valid := ctx.Request.Validate(map[string]string{
		"name": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

    // Create and send response
	user := models.User{
		Name: params["name"].(string),
	}
	ctx.Database.Create(&user)
	ctx.Response.JSON(&UserResponse{Success: true, User: &user})
}
```

## Update

To update model we need to make sure it exists first before doing anything to it. Here's how you can do it:

```go
func (c *UserController) Update(ctx *framework.Context) {

    // Validate request parameters
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

    // Make sure user exists
	if User.ID == 0 {
		ctx.Response.JSON(&UserResponse{Success: false, Message: "User doesn't exist"})
		return
	}

    // Update user column and send response
	ctx.Database.Model(&user).Update("name", params["name"])
	ctx.Response.JSON(&UserResponse{Success: true, User: &user})
}
```

## Delete

Delete operations are also possible. Remember to validate that the model exists before actually calling the Delete function on the Database.

```go
func (c *UserController) Delete(ctx *framework.Context) {
	
    // Validate request parameters
    params, valid := ctx.Request.Validate(map[string]string{
		"id": "required",
	})

	if !valid {
		ctx.Response.String("Validation failed")
		return
	}

	var User models.User
	ctx.Database.First(&user, params["id"])

    // Check that the user exists
	if User.ID == 0 {
		ctx.Response.JSON(&UserResponse{Success: false, Message: "User doesn't exist"})
		return
	}

    // Delete the user and send response
	ctx.Database.Delete(&user)
	ctx.Response.JSON(&UserResponse{Success: true})

}
```
> Please make sure to review the `GORM` documentation to learn more on how the database operations work.