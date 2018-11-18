# Commands (CLI)

There's a small cmd in the root of the project called `golosina`. This has a couple of commands that are already prepared for you. This works using the [Cobra Library](https://github.com/spf13/cobra).

## Adding new commands

Simple, just create a new file under the cmd folder/package and add whatever you need to run. Follow the other files as example on how to write them and also you can check the [Cobra Library](https://github.com/spf13/cobra) documentation for more information. 

Also remember to run the refresh command when you're done so it's accessible by the `golosina` tool. 


## Available commands

To see whats available just run:

```bash
./golosina
```

And you'll see a list of available commands for this app. The most important ones that we can see here are the `serve` and the `refresh` but you can add more.

### Serve Cmd

This will start the web server so you can begin listening for requests. The way to use it is the following:

```bash
./golosina serve
```
!> If you delete the file located in `cmd/serve` this will stop working. 

### Refresh Cmd

This will rebuild the golosina cmd tool. Adding or removing whatever you have on the cmd package. To run it is basically the same as the serve cmd:

```bash
./golosina refresh
```
!> If you delete the file located in `cmd/refresh` this will stop working. 
