// handlers.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetAllHandler(w http.ResponseWriter, r *http.Request) {
	u := new(User)
	users, err := u.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusFound)
		return
	}

	j, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(j)
}
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Query id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Query id must be a number", http.StatusBadRequest)
		return
	}

	u := new(User)
	user, err := u.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = user.CreateUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Notes
func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = note.CreateNote()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUserNotesHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Query id is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Query id must be a number", http.StatusBadRequest)
		return
	}

	n := new(Note)
	note, err := n.GetNote(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	j, err := json.Marshal(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// List
func ListPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var list List
	err := decoder.Decode(&list)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}

	list.order()
	response, err := list.ToJson()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
