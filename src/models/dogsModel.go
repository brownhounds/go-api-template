package models

import (
	"go-api-template/src/clients"
	"time"
)

type Dog struct {
	Id        string    `db:"id"`
	Name      string    `db:"name"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
}

func GetDogs() *[]Dog {
	con := clients.Postgres.Connection()
	dogs := []Dog{}

	err := con.Select(&dogs, "SELECT * FROM get_dogs()")
	if err != nil {
		panic(err)
	}

	return &dogs
}
