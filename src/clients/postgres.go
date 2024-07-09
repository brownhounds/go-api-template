package clients

import (
	"fmt"
	c "go-api-template/src/config"
	"time"

	"github.com/brownhounds/swift/env"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Postgres *postgres

type postgres struct {
	DB         *sqlx.DB
	connection string
}

func PostgresInit() {
	Postgres = &postgres{
		DB: nil,
		connection: fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s sslmode=%s search_path=%s",
			env.Env(c.SQL_USER),
			env.Env(c.SQL_PASSWORD),
			env.Env(c.SQL_DB),
			env.Env(c.SQL_HOST),
			env.Env(c.SQL_SSL),
			env.Env(c.SQL_DB_SCHEMA),
		),
	}

	db, err := sqlx.Connect("postgres", Postgres.connection)
	if err != nil {
		panic(err)
	}

	Postgres.DB = db

	Postgres.DB.SetMaxIdleConns(
		env.EnvInt(c.POSTGRES_POOL_MAX_IDLE_CONNS),
	)
	Postgres.DB.SetConnMaxIdleTime(
		env.EnvTimeDuration(c.POSTGRES_POOL_MAX_IDLE_TIME_SECONDS) * time.Second,
	)
	Postgres.DB.SetMaxOpenConns(
		env.EnvInt(c.POSTGRES_POOL_MAX_OPEN_CONNS),
	)
}

func (p *postgres) Connection() *sqlx.DB {
	return p.DB
}
