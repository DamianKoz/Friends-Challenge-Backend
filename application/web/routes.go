package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/challenge/create", challengeCreate)
	mux.HandleFunc("/challenge/view", ChallengeView)
	mux.HandleFunc("/user/create", app.userCreate)
	mux.HandleFunc("/user/view", app.userView)
	mux.HandleFunc("/users", app.allUsers)

	return mux
}
