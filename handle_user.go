package main

import (
	"fmt"
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/login.html")
}

func handleRegisterUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/register.html")
}

func handleUserProfile(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("userName")
	w.Write([]byte(fmt.Sprintf("stub for user %s\n", name)))

}
