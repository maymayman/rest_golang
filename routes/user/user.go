// package main
package math
import "github.com/kataras/iris"

func Routes(app *iris.Application) {
	userMiddleware := func(ctx iris.Context) {
		ctx.Next()
	}

	router := app.Party("/users", userMiddleware) 
   
    router.Get("/", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "message": "get user",
        })
	})
	
	router.Post("/", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
			"message": "create user",
			"average": "avg",
        })
    })
}