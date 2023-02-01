package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	GET := http.MethodGet
	POST := http.MethodPost

	router := httprouter.New()

	// Dynamic Routes with sessions
	dynamic := alice.New(app.sessionManger.LoadAndSave)

	// Home -> Inital Welcoming
	router.HandlerFunc(GET, "/", app.home)

	// Routes for challenges
	router.Handler(POST, "/challenge/create", dynamic.ThenFunc(app.challengeCreate))
	router.Handler(GET, "/challenge/view/:id", dynamic.ThenFunc(app.ChallengeView))
	router.Handler(GET, "/challenge/all/", dynamic.ThenFunc(app.AllChallenges))

	// Routes for users
	router.HandlerFunc(POST, "/user/signup", app.userCreatePost)
	router.HandlerFunc(GET, "/user/login", app.userLogin) // Delete?
	router.HandlerFunc(POST, "/user/login", app.userLoginPost)
	router.HandlerFunc(POST, "/user/logout", app.userLogoutPost)
	router.HandlerFunc(GET, "/user/view/:id", app.userView)
	router.HandlerFunc(GET, "/user/all", app.allUsers)

	standard := alice.New(app.logRequest, secureHeaders)
	return standard.Then(router)

	// return app.logRequest(secureHeaders(mux)) -> Same as above
}
