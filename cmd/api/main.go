package main

import (
	"farmaIA/internal/server"
	"fmt"
    "net/http"
    httpSwagger "github.com/swaggo/http-swagger"
    "github.com/go-chi/chi/v5"
    _ "farmaIA/cmd/api/swagger"
)

func main() {

	server := server.NewServer()

   r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), //The url pointing to API definition
	))

	http.ListenAndServe(":8080", r)    

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
