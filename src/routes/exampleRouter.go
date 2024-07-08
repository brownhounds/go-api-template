package routes

import (
	"go-api-template/src/handlers"
	"go-api-template/src/middlewares"

	"github.com/brownhounds/swift"
)

func ExampleRouter(app *swift.Swift) {
	v1 := app.Group("/v1")
	v1.Get("/test", handlers.ExampleHandler).Middleware(middlewares.ExampleMiddleware)
	v1.Get("/test/{id}", handlers.ExampleHandlerWithId)
	v1.Get("/dogs", handlers.GetDogsHandler)
}
