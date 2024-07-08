package routes

import (
	"go-api-template/handlers"
	"go-api-template/middlewares"

	"github.com/brownhounds/swift"
)

func ExampleRouter(app *swift.Swift) {
	v1 := app.Group("/v1")
	v1.Get("/test", handlers.ExampleHandler).Middleware(middlewares.ExampleMiddleware)
	v1.Get("/test/{id}", handlers.ExampleHandlerWithId)
}
