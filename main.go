package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/li-go/jwt-auth-service/handlers"
	"github.com/li-go/jwt-auth-service/middlewares"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/public", handlers.Public)
	r.Handle("/sign_in", handlers.SignIn)

	r.Handle("/private", middlewares.AuthHandler(handlers.Private))

	log.Println("ListenAndServe :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("can not start server", err)
	}
}
