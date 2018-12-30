package main

import "github.com/kataras/iris"
import (
	user "./routes/user"
)
import "fmt"

func authentication(ctx iris.Context) {
	fmt.Printf("user")
	ctx.Next()
}

func createServer() *iris.Application {
	app := iris.New()
	app.Use(func(ctx iris.Context){ctx.Next()}, authentication)
	user.Routes(app)

	return app
}

func main() {
	app := createServer()
    app.Run(iris.Addr(":8080"))
}