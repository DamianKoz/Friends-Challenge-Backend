package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Create a slice with all the paths to the needed files
	files := []string{"./ui/html/templates/base.tmpl.html", "./ui/html/pages/index.html", "./ui/html/partials/nav.tmpl.html"}

	// Use the template.ParseFiles() function to read the template file into a template set.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	// We then use the Execute() method on the template set to write the template content as the response body.
	// The last parameter to Execute() represents any dynamic data that we want to pass in.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func ChallengeView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Found Challenge with ID: %v", id)

}

func challengeCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Creating Challenge"))
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

func (app *application) userView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID. ERROR:"+err.Error(), http.StatusInternalServerError)
		return
	}

	user, err := app.user.Get(id)
	if err != nil {
		http.Error(w, "Could not Get User. ERROR: "+err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
	// fmt.Fprintf(w, "Found User with ID %v: %+v", id, user)

}

func (app *application) allUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Fetching all users...\n")
	users, err := app.user.GetAll()
	if err != nil {
		http.Error(w, "Error occurred while fetching all users"+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Successfully fetched all users\n")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)

	// json, _ := json.Marshal(users)
	// w.Write(json)

	// fmt.Fprintf(w, "\nAll Users:\n")
	// for _, user := range users {
	// 	fmt.Fprintf(w, "User: %v, %v \n", user.FirstName, user.LastName)
	// }
}
