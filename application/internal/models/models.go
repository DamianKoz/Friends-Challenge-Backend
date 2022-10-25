package models

import (
	"database/sql"
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

type Task struct {
	// Amount + Name per Unit -> 5 Squats per Day
	ID         uint
	Title      string
	Amount     uint
	AmountUnit string
	TimeUnit   string
}

type Tasks []Task

type TaskModel struct {
	DB *sql.DB
}

func (t *TaskModel) Insert(title, verb string, amount int, activity, duration string, user_id int, time_unit string) (int, error) {

	statement := `INSERT INTO task (title, verb, amount, activity, duration, user_id, time_unit) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	if err := t.DB.Ping(); err != nil {
		return 0, err
	}
	result, err := t.DB.Exec(statement, title, verb, amount, activity, duration, user_id, time_unit)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	// fmt.Printf("Inserted '%v', '%v', with id of '%v'", firstName, lastName, id)

	return int(id), nil
}
