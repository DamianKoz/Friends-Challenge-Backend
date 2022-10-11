package models

import (
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
