package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type greeting struct {
	Message string
	Status  bool
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Returns an inital greeting to check, whether the API is working
	answer := greeting{Message: "Api is working :)", Status: true}

	json.NewEncoder(w).Encode(answer)
}

func (app *application) ChallengeView(w http.ResponseWriter, r *http.Request) {
	// Return a challenge which corresponds to the id from the url

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Found Challenge with ID: %v", id)

}

func (app *application) challengeCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Giving FORM to create new Challenge"))
}

func (app *application) challengeCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created Challenge with Form Data"))
}

func (app *application) userCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	firstName := "Julian"
	lastName := "Fritz"

	id, err := app.user.Insert(firstName, lastName)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/user/view?id=%d", id), http.StatusSeeOther)
}
func (app *application) userCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Displaying User Creation Form"))

	// id, err := app.user.Insert(firstName, lastName)
	// if err != nil {
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	// http.Redirect(w, r, fmt.Sprintf("/user/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) userView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	user, err := app.user.Get(id)
	if err != nil {
		http.Error(w, "Could not Get User. ERROR: "+err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (app *application) allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Fetching all users...\n")
	users, err := app.user.GetAll()
	if err != nil {
		http.Error(w, "Error occurred while fetching all users"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Successfully fetched all users\n")

	json.NewEncoder(w).Encode(users)

	// json, _ := json.Marshal(users)
	// w.Write(json)

	// fmt.Fprintf(w, "\nAll Users:\n")
	// for _, user := range users {
	// 	fmt.Fprintf(w, "User: %v, %v \n", user.FirstName, user.LastName)
	// }
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	// Get Method: Display HTML Form for logging in
}
func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	// Post Method: Authenticate and login the user
}
func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
	// Post Method: Logout the user
}
