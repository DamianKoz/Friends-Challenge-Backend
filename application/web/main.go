package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

var (
	FlagSrcFolder  = flag.String("src", "./pages/", "blog folder")
	FlagStaticF    = flag.String("static", "./ui/static/", "static folder")
	FlagTmplFolder = flag.String("tmpl", "./templates/", "template folder")
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

func main() {
	port := ":8080"

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/challenge/create", challengeCreate)
	mux.HandleFunc("/challenge/view", ChallengeView)

	log.Println("Starting Server on " + port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)

}
