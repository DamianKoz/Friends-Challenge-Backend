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

// Handlers for Tasks

type taskCreateForm struct {
	User_ID   int
	Title     string
	Verb      string
	Amount    int
	Activity  string
	Time_Unit string
	Duration  string
}

func (app *application) taskCreate(w http.ResponseWriter, r *http.Request) {
	// Creates a new task. Expects: title, amount, activity, duration, user_id, end_date
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Extraction of Value from Form

	amount, err := strconv.Atoi(r.PostForm.Get("task_amount"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	form := taskCreateForm{
		User_ID:   2, // HARDCODED FOR NOW!!
		Title:     r.PostForm.Get("task_title"),
		Verb:      r.PostForm.Get("task_verb"),
		Activity:  r.PostForm.Get("task_activity"),
		Time_Unit: r.PostForm.Get("task_time_unit"),
		Duration:  r.PostForm.Get("task_duration"),
		Amount:    amount,
	}

	id, err := app.task.Insert(form.Title, form.Verb, form.Amount, form.Activity, form.Duration, form.User_ID, form.Time_Unit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// fmt.Printf("New Task: %v\n%v %v per %v. Duration: %v Days", title, amount, activity, time_unit, duration)
	fmt.Printf("Everything worked. ID of created Task: %v", id)
	http.Redirect(w, r, "http://127.0.0.1:5500/create_task.html", http.StatusSeeOther)

}

// Handlers for Challenges

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

// Handlers for User Management

func (app *application) userCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Validating Data ... \nCreated new User."))

	// firstName := "Julian"
	// lastName := "Fritz"

	// id, err := app.user.Insert(firstName, lastName)
	// if err != nil {
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	// http.Redirect(w, r, fmt.Sprintf("/user/view?id=%d", id), http.StatusSeeOther)
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

// Querying Users

func (app *application) allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Fetching all users...\n")
	users, err := app.user.GetAll()
	if err != nil {
		http.Error(w, "Error occurred while fetching all users"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Successfully fetched all users\n")

	json.NewEncoder(w).Encode(users)
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
