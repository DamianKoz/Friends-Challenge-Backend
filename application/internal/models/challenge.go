package models

import (
	"database/sql"
	"fmt"
)

type Challenge struct {
	// Verb + Amount + Name per Unit -> Do 5 Squats per Day
	Id        int
	User_ID   int
	Title     string
	Verb      string
	Amount    int
	Activity  string
	Time_Unit string
	Duration  string
}

type ChallengeCreateForm struct {
}

type Users []User

type ChallengeModel struct {
	DB *sql.DB
}

func (c *ChallengeModel) Insert(title, verb string, amount int, activity, duration string, user_id int, time_unit string) (int, error) {

	statement := `INSERT INTO task (title, verb, amount, activity, duration, user_id, time_unit) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	if err := c.DB.Ping(); err != nil {
		return 0, err
	}
	result, err := c.DB.Exec(statement, title, verb, amount, activity, duration, user_id, time_unit)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (c *ChallengeModel) GetAll() ([]*Challenge, error) {

	statement := `
	SELECT 
		id, title, amount, activity, user_id, time_unit, verb
	FROM 
		task`
	rows, err := c.DB.Query(statement)

	if err != nil {
		fmt.Printf("Error occured while fetching from DB.")
		return nil, err
	}

	defer rows.Close()

	challenges := []*Challenge{}

	for rows.Next() {
		challenge := &Challenge{}
		if err := rows.Scan(&challenge.Id, &challenge.Title, &challenge.Amount, &challenge.Activity, &challenge.User_ID, &challenge.Time_Unit, &challenge.Verb); err != nil {
			fmt.Printf("Error occured while creating Challenge Objects from DB data.")
			return nil, err
		}
		challenges = append(challenges, challenge)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return challenges, nil
}
