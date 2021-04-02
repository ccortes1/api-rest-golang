//main.go
package main

import (
	"database/sql"
	"flag"
	"log"

	//Driver para sqlite
	_ "modernc.org/sqlite"
)

func main() {

	server := NewServer(":3000")

	migrate := flag.Bool("migrate", false, "create tables of db")
	flag.Parse()
	if *migrate {
		if err := MakeMigrations(); err != nil {
			log.Fatal(err)
		}
	}

	server.Handle("POST", "/list", ListPostRequest)
	server.Handle("GET", "/users", GetAllHandler)
	server.Handle("POST", "/users", CreateUserHandler)
	server.Handle("GET", "/users/us", GetUserHandler)
	server.Handle("POST", "/note", CreateNoteHandler)
	server.Handle("GET", "/user/notes", GetUserNotesHandler)
	server.Listen()

}

var db *sql.DB

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error

	db, err = sql.Open("sqlite", "data.sqlite")
	if err != nil {
		panic(err)
	}
	return db
}

func MakeMigrations() error {
	db := GetConnection()

	q := `CREATE TABLE IF NOT EXISTS users (
		id_user INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(100),
		email VARCHAR(100),
		phone VARCHAR(15));

		CREATE TABLE IF NOT EXISTS notes (
		id_note INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(100),
		description VARCHAR(200),
		user_id INTEGER);
		`
	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}
