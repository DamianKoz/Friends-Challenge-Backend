package main

import "net/http"

func routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/challenge/create", challengeCreate)
	mux.HandleFunc("/challenge/view", ChallengeView)

	return mux
}
