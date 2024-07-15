package handlers

import (
	"go-api-template/src/models"
	"net/http"
	"time"

	"github.com/brownhounds/swift/res"
)

type Dog struct {
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
	Name      string    `db:"name"`
	Id        string    `db:"id"`
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	res.Json(w, http.StatusOK, res.Map{"message": "Hello There!"})
}

func ExampleHandlerWithId(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	res.Json(w, http.StatusOK, res.Map{"id": id})
}

func GetDogsHandler(w http.ResponseWriter, r *http.Request) {
	dogs := models.GetDogs()
	res.Json(w, http.StatusOK, dogs)
}
