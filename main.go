package main

import (
	"fmt"
	"net/http"

	"github.com/TChi91/ChitChat/data"
)

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	// mux.HandleFunc("/err", err)
	// mux.HandleFunc("/login", login)
	// mux.HandleFunc("/logout", logout)
	// mux.HandleFunc("/signup", signup)
	// mux.HandleFunc("/signup_account", signupAccount)
	// mux.HandleFunc("/authenticate", authenticate)
	// mux.HandleFunc("/thread/new", newThread)
	// mux.HandleFunc("/thread/create", createThread)
	// mux.HandleFunc("/thread/post", postThread)
	// mux.HandleFunc("/thread/read", readThread)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		fmt.Fprint(writer, "Cannot get threads")
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, threads, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, threads, "layout", "private.navbar", "index")
		}
	}
}
