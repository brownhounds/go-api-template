package handlers

import (
	"net/http"

	res "github.com/brownhounds/swift/response"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	res.Json(w, http.StatusOK, res.Map{"message": "Hello There!"})
}

func ExampleHandlerWithId(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	res.Json(w, http.StatusOK, res.Map{"id": id})
}
