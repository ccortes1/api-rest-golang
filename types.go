// types.go
package main

import (
	"encoding/json"
	"errors"

	//Driver para sqlite
	_ "modernc.org/sqlite"
)

// User
type User struct {
	ID    int    `json:"id_user"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u User) CreateUser() error {
	db := GetConnection()
	q := `INSERT INTO users(name, email,phone)VALUES(?,?,?)`
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	r, err := stmt.Exec(u.Name, u.Email, u.Phone)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("Error: a row was expected")
	}

	return nil
}

func (u User) GetAllUsers() ([]User, error) {
	db := GetConnection()
	q := `SELECT id_user, name, email, phone FROM users`

	rows, err := db.Query(q)
	if err != nil {
		return []User{}, err
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Phone,
		)
		users = append(users, u)
	}
	return users, nil
}

func (u User) GetUser(id int) (User, error) {
	//var user User
	db := GetConnection()
	q := `SELECT id_user, name, email, phone FROM users WHERE id_user=?`
	row := db.QueryRow(q, id)
	err := row.Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Phone,
	)
	if err != nil {
		return u, err
	}

	return u, nil
}

// Notes
type Note struct {
	ID          int    `json:"id_note"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      int    `json:"user_id"`
}

func (n Note) CreateNote() error {
	db := GetConnection()
	q := `INSERT INTO notes (title, description, user_id)VALUES(?,?,?)`

	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	r, err := stmt.Exec(n.Title, n.Description, n.UserId)
	if err != nil {
		return err
	}

	if i, err := r.RowsAffected(); err != nil || i != 1 {
		return errors.New("Error: a row was expected")
	}

	return nil

}

func (n Note) GetNote(id int) ([]Note, error) {
	db := GetConnection()
	q := `SELECT id_note, title, description, user_id FROM notes WHERE user_id=?`
	rows, err := db.Query(q, id)
	if err != nil {
		return []Note{}, err
	}

	defer rows.Close()
	notes := []Note{}
	for rows.Next() {
		rows.Scan(
			&n.ID,
			&n.Title,
			&n.Description,
			&n.UserId,
		)
		notes = append(notes, n)
	}
	return notes, nil
}

// Lists
type List struct {
	SinClasificar []int `json:"sin clasificar"`
	Clasificado   []int `json:"clasificado"`
}

func (l *List) ToJson() ([]byte, error) {
	return json.Marshal(l)
}

func (l *List) order() {
	ar := l.SinClasificar
	var end []int
	var rep []int
	dic := make(map[int]int)

	for _, value := range ar {
		if dic[value] == 0 {
			for _, value2 := range ar {
				if value == value2 {
					dic[value]++
					if dic[value] >= 2 {
						rep = append(rep, value2)
					}
				}
			}
		}
	}
	for k, _ := range dic {
		end = append(end, k)
	}
	end = sort(end)
	end = append(end, rep...)
	l.Clasificado = end
	return
}

func sort(l []int) []int {
	list := l
	for i, _ := range list {
		min_idx := i
		for j, value2 := range list {
			if list[min_idx] < value2 {
				min_idx = j
			}
			list[i], list[min_idx] = list[min_idx], list[i]
		}

	}

	return list
}
