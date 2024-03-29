package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"

	"github.com/damiankoz/friends_challenge/application/internal/models"
	_ "github.com/go-sql-driver/mysql"
)

var (
	FlagSrcFolder  = flag.String("src", "./pages/", "blog folder")
	FlagStaticF    = flag.String("static", "./ui/static/", "static folder")
	FlagTmplFolder = flag.String("tmpl", "./templates/", "template folder")
)

type config struct {
	addr string
}

var cfg config

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	user          *models.UserModel
	challenge     *models.ChallengeModel
	sessionManger *scs.SessionManager
}

func main() {
	flag.StringVar(&cfg.addr, "addr", ":8080", "HTTP network address")
	// For use with Docker: host.docker.internal
	dsn := flag.String("dsn", "web:Korona11@tcp(127.0.0.1:3306)/friends_challenge?parseTime=true", "MySQL data source name")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDb(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		user:          &models.UserModel{DB: db},
		challenge:     &models.ChallengeModel{DB: db},
		sessionManger: sessionManager,
	}

	mux := app.routes()

	infoLog.Printf("Starting Server on http://localhost" + cfg.addr)
	err = http.ListenAndServe(cfg.addr, mux)
	errorLog.Fatal(err)

}
