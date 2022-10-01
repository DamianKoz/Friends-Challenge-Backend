package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Challenge struct {
	ID           uint
	Name         string
	Participants Users
	CreatedBy    User
	Tasks        Tasks
	StartTime    time.Time
	EndTime      time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Users []User

type User struct {
	Id                     uint
	Profile                Profile
	Level                  uint
	Exp                    uint
	ParticipatedChallenges uint
	CompletedChallenges    uint
}

type Profile struct {
	ID        uint
	FirstName string
	LastName  string
	Age       uint
	Gender    string
	Email     string
	Birthday  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Task struct {
	// Amount + Name per Unit -> 5 Squats per Day
	ID         uint
	Name       string
	Amount     uint
	AmountUnit string
	TimeUnit   string
}

type Tasks []Task

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World"))
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

func main() {
	port := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/challenge/create", challengeCreate)
	mux.HandleFunc("/challenge/view", ChallengeView)

	log.Println("Starting Server on " + port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)

}
