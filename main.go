package main

import (
	"go-api-template/config"
	"go-api-template/routes"
	"os"

	"github.com/brownhounds/swift"
	"github.com/brownhounds/swift/env"
)

func main() {
	env.InitWithMandatoryVariables(config.MandatoryEnvVariables)

	app := swift.New()
	app.OApiValidator("schema.yml")

	app.SwaggerStaticServer(".swagger", "/docs")
	app.SwaggerServe(true)

	routes.ExampleRouter(app)

	app.AddTLS(".certs/server.pem", ".certs/server-key.pem")
	app.Serve(os.Getenv(config.SERVER_HOST), os.Getenv(config.SERVER_PORT))
}
