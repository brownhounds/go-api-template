package config

const (
	SERVER_HOST = "SERVER_HOST"
	SERVER_PORT = "SERVER_PORT"
)

var MandatoryEnvVariables = []string{
	SERVER_HOST,
	SERVER_PORT,
}
