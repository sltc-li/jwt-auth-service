package handlers

import (
	"net/http"

	"github.com/li-go/jwt-auth-service/middlewares"
	"github.com/li-go/jwt-auth-service/repositories"
)

var Public = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Anonymous"))
})

var SignIn = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	password := r.URL.Query().Get("password")
	if len(name) == 0 || len(password) == 0 {
		http.Error(w, "require name and password", http.StatusUnauthorized)
		return
	}

	userRepository := &repositories.User{}
	user, err := userRepository.FindByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if password != user.Password {
		http.Error(w, "wrong password", http.StatusUnauthorized)
		return
	}

	signedToken, err := middlewares.ToSignedToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Write([]byte(signedToken))
})
