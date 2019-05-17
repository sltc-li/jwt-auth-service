package handlers

import (
	"fmt"
	"net/http"

	"github.com/li-go/jwt-auth-service/middlewares"
)

var Private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	user, err := middlewares.FromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	fmt.Fprintf(w, "Hello, %s<#%d>", user.Name, user.ID)
})
