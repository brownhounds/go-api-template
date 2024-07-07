package main

import (
	"go-api-template/routes"
	"os"

	"github.com/brownhounds/swift"
	"github.com/brownhounds/swift/env"
)

func main() {
	env.InitWithMandatoryVariables([]string{
		"PORT",
	})

	app := swift.New()
	app.OApiValidator("schema.yml")

	app.SwaggerStaticServer(".swagger", "/docs")
	app.SwaggerServe(true)

	routes.ExampleRouter(app)

	app.AddTLS(".certs/server.pem", ".certs/server-key.pem")
	app.Serve("0.0.0.0", os.Getenv("PORT"))
}
