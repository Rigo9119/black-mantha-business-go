package main

import "github.com/kataras/iris/v12"

func main() {
	app := irisNewApp()

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func irisNewApp() *iris.Application {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.Text("Hello MIgue kun")
	})

	return app
}
