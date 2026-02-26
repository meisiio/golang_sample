package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
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
		fmt.Println("could not to fetch user", err)
		return
	}
	fmt.Printf("%+v", person)
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
