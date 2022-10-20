package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	// fileserver := http.FileServer(http.Dir("./ui/static/"))
	// mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	// mux.HandleFunc("/", home)

	mux.HandleFunc("/challenge/create", challengeCreate)
	mux.HandleFunc("/challenge/view", ChallengeView)

	mux.HandleFunc("/user/create", app.userCreate)
	mux.HandleFunc("/user/view", app.userView)
	mux.HandleFunc("/users", app.allUsers)

	standard := alice.New(app.logRequest, secureHeaders)
	return standard.Then(mux)

	// return app.logRequest(secureHeaders(mux)) -> Same as above
}
