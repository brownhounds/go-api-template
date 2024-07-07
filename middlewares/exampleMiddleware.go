package middlewares

import (
	"fmt"
	"net/http"
)

func ExampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello From Middleware")
		next.ServeHTTP(w, r)
	})
}
