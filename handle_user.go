package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/login.html")
}

func handleGetRegisterUser(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/register.html")
}

func handlePostRegisterUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Printf("Unable to parse form: %s", err)
		http.Error(w, "bad", http.StatusBadRequest)
		return
	}
	pUser := r.PostForm.Get("user")
	if pUser == "" {
		http.Error(w, "no user", http.StatusBadRequest)
		return
	}
	pPassword := r.PostForm.Get("password")
	if pPassword == "" {
		http.Error(w, "no pw", http.StatusBadRequest)
		return
	}

	w.Write([]byte("made it"))

}

func handleUserProfile(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("userName")
	w.Write([]byte(fmt.Sprintf("stub for user %s\n", name)))

}
