package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	FlagSrcFolder  = flag.String("src", "./pages/", "blog folder")
	FlagStaticF    = flag.String("static", "./ui/static/", "static folder")
	FlagTmplFolder = flag.String("tmpl", "./templates/", "template folder")
)

type config struct {
	addr string
}

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

var cfg config

func main() {
	flag.StringVar(&cfg.addr, "addr", ":8080", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := routes()

	infoLog.Printf("Starting Server on http://localhost" + cfg.addr)
	err := http.ListenAndServe(cfg.addr, mux)
	errorLog.Fatal(err)

}
