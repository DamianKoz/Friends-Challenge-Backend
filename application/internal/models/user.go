package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id                     uint
	Level                  uint
	Exp                    uint
	ParticipatedChallenges uint
	CompletedChallenges    uint
	FirstName              string
	LastName               string
	Age                    uint
	Gender                 string
	Email                  string
	Birthday               time.Time
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(FirstName, LastName string) (int, error) {
	statement := `INSERT INTO profile (firstName, lastName) 
	VALUES (?, ?)`
	result, err := u.DB.Exec(statement, FirstName, LastName)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(id), nil
}

func (u *UserModel) Get(id int) (*User, error) {
	return nil, nil
}
