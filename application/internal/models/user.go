package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	Id                     uint
	FirstName              string
	LastName               string
	Level                  uint
	Exp                    uint
	ParticipatedChallenges uint
	CompletedChallenges    uint
	Age                    uint
	Gender                 string
	Email                  string
	Birthday               sql.NullTime
	CreatedAt              sql.NullTime
	UpdatedAt              sql.NullTime
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(firstName, lastName string) (int, error) {
	// db, _ := sql.Open("mysql", "web:Korona11@/friends_challenge?parseTime=true")

	statement := `INSERT INTO user (first_name, last_name) 
	VALUES (?, ?)`
	if err := u.DB.Ping(); err != nil {
		return 0, err
	}
	result, err := u.DB.Exec(statement, firstName, lastName)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	fmt.Printf("Inserted '%v', '%v', with id of '%v'", firstName, lastName, id)

	return int(id), nil
}

func (u *UserModel) Get(id int) (*User, error) {
	// Gets an user with his id, firstname, lastname, level and exp

	// SHORTER Alternative
	// err := u.DB.QueryRow("SELECT ...", id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Level, &user.Exp)

	statement := "SELECT * FROM user where id = ?"
	row := u.DB.QueryRow(statement, id)
	user := User{}
	err := row.Scan(&user.Id, &user.Level, &user.Exp, &user.FirstName, &user.LastName, &user.Age, &user.Gender, &user.Email, &user.UpdatedAt, &user.CreatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("User Id does not exist")
		} else {
			return nil, err
		}
	}
	return &user, nil
}

func (u *UserModel) GetAll() ([]*User, error) {
	statement := `
	SELECT 
		id, user_level, exp, first_name, last_name, age, gender, email, updated_at, created_at 
	FROM 
		user`
	rows, err := u.DB.Query(statement)

	if err != nil {
		fmt.Printf("Error occured while fetching from DB.")
		return nil, err
	}

	defer rows.Close()

	users := []*User{}

	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.Id, &user.Level, &user.Exp, &user.FirstName, &user.LastName, &user.Age, &user.Gender, &user.Email, &user.UpdatedAt, &user.CreatedAt); err != nil {
			fmt.Printf("Error occured while creating User Objects from DB data.")
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
