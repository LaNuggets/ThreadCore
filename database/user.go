package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	id             int
	profilePicture string
	email          string
	username       string
	password       string
}

func addUser(db *sql.DB, user user) {
	query, _ := db.Prepare("INSERT INTO user (id, profilePicture, email, username, password) VALUES (?, ?, ?, ?, ?)")
	query.Exec(user.id, user.profilePicture, user.email, user.username, user.password)
	defer query.Close()
}

func searchForPerson(db *sql.DB, searchString string) []person {

	rows, err := db.Query("SELECT id, first_name, last_name, email, ip_address FROM people WHERE first_name like '%" + searchString + "%' OR last_name like '%" + searchString + "%'")

	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	people := make([]person, 0)

	for rows.Next() {
		ourPerson := person{}
		err = rows.Scan(&ourPerson.id, &ourPerson.first_name, &ourPerson.last_name, &ourPerson.email, &ourPerson.ip_address)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, ourPerson)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return people
}

func getPersonById(db *sql.DB, ourID string) person {

	fmt.Printf("Value is %v", ourID)

	rows, _ := db.Query("SELECT id, first_name, last_name, email, ip_address FROM people WHERE id = '" + ourID + "'")
	defer rows.Close()

	ourPerson := person{}

	for rows.Next() {
		rows.Scan(&ourPerson.id, &ourPerson.first_name, &ourPerson.last_name, &ourPerson.email, &ourPerson.ip_address)
	}

	return ourPerson
}

func updatePerson(db *sql.DB, ourPerson person) int64 {

	stmt, err := db.Prepare("UPDATE people set first_name = ?, last_name = ?, email = ?, ip_address = ? where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(ourPerson.first_name, ourPerson.last_name, ourPerson.email, ourPerson.ip_address, ourPerson.id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	return affected
}

func deletePerson(db *sql.DB, idToDelete string) int64 {

	stmt, err := db.Prepare("DELETE FROM people where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(idToDelete)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	return affected

}
