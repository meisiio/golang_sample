package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func dbTest() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=meisam dbname=postgres password=newpassword sslmode=disable connect_timeout=10")
	if err != nil {
		fmt.Println("opening did not occured", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("db does not response", err)
		return
	}
	person, err := userfetch(1, db)
	if err != nil {
		fmt.Println("could not fetch user by id", err)
		return
	}
	fmt.Printf("single user: %+v\n", person)

	users, err2 := userfetchwithquery("ali", db)
	if err2 != nil {
		fmt.Println("db was not able to fetch rows", err2)
		return
	}
	fmt.Printf("matched users: %+v\n", users)

}

type user struct {
	fullname string
	id       int
	role     string
}

func userfetch(id int, db *sql.DB) (*user, error) {
	userm := user{}
	err := db.QueryRow("SELECT id, fullname, role FROM users WHERE id = $1", id).Scan(&userm.id, &userm.fullname, &userm.role)
	if err != nil {
		return &userm, err
	}
	return &userm, nil
}
func userfetchwithquery(fullname string, db *sql.DB) ([]user, error) {
	rows, err := db.Query("SELECT id, fullname, role FROM users WHERE fullname ILIKE '%' || $1 || '%'", fullname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		u := user{}
		if err := rows.Scan(&u.id, &u.fullname, &u.role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
