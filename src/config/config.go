package config

const (
	SERVER_HOST                         = "SERVER_HOST"
	SERVER_PORT                         = "SERVER_PORT"
	SQL_USER                            = "SQL_USER"
	SQL_PASSWORD                        = "SQL_PASSWORD"
	SQL_DB                              = "SQL_DB"
	SQL_HOST                            = "SQL_HOST"
	SQL_SSL                             = "SQL_SSL"
	SQL_DB_SCHEMA                       = "SQL_DB_SCHEMA"
	POSTGRES_POOL_MAX_IDLE_CONNS        = "POSTGRES_POOL_MAX_IDLE_CONNS"
	POSTGRES_POOL_MAX_IDLE_TIME_SECONDS = "POSTGRES_POOL_MAX_IDLE_TIME_SECONDS"
	POSTGRES_POOL_MAX_OPEN_CONNS        = "POSTGRES_POOL_MAX_OPEN_CONNS"
)

var MandatoryEnvVariables = []string{
	SERVER_HOST,
	SERVER_PORT,
	SQL_USER,
	SQL_PASSWORD,
	SQL_DB,
	SQL_HOST,
	SQL_SSL,
	SQL_DB_SCHEMA,
	POSTGRES_POOL_MAX_IDLE_CONNS,
	POSTGRES_POOL_MAX_IDLE_TIME_SECONDS,
	POSTGRES_POOL_MAX_OPEN_CONNS,
}
