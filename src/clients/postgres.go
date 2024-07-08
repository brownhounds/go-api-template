package clients

import (
	"fmt"
	c "go-api-template/src/config"
	"os"
	"time"

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
			os.Getenv(c.SQL_USER),
			os.Getenv(c.SQL_PASSWORD),
			os.Getenv(c.SQL_DB),
			os.Getenv(c.SQL_HOST),
			os.Getenv(c.SQL_SSL),
			os.Getenv(c.SQL_DB_SCHEMA),
		),
	}

	db, err := sqlx.Connect("postgres", Postgres.connection)
	if err != nil {
		panic(err)
	}

	Postgres.DB = db

	Postgres.DB.SetMaxIdleConns(
		c.EnvInt(c.POSTGRES_POOL_MAX_IDLE_CONNS),
	)
	Postgres.DB.SetConnMaxIdleTime(
		time.Duration(c.EnvInt(c.POSTGRES_POOL_MAX_IDLE_TIME_SECONDS)) * time.Second,
	)
	Postgres.DB.SetMaxOpenConns(
		c.EnvInt(c.POSTGRES_POOL_MAX_OPEN_CONNS),
	)
}

func (p *postgres) Connection() *sqlx.DB {
	return p.DB
}
