package main

import (
	"database/sql"
	"fmt"
	"log"
)
import "time"

func main() {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/root?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{
		query := `
		CREATE TABLE USERS (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	);`
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		username := "Leo Messi"
		password := "goat"
		created_at := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, created_at)
		if err != nil {
			log.Fatal(err)
		}

		userID, err := result.LastInsertId()
		fmt.Println(userID)
	}
	{
		var (
			id         int
			username   string
			password   string
			created_at time.Time
		)
		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &created_at); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, created_at)
	}
	{
		type user struct {
			id         int
			username   string
			password   string
			created_at time.Time
		}
		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		var users []user
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.password, &u.created_at)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v", users)
	}
	{
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
