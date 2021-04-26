package controllers

import (
	"backend/src/database"
	"backend/src/models"
	"backend/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// List all users
func Show(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show not implemented!"))
}

// List a user by id
func Find(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Find not implemented!"))
}

// Create a user
func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.UsersRepository(db)

	userID, err := repository.Create(user)

	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("created ok! ID: %d", userID)))
}

// Update one user by id
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update not implemented!"))
}

// Remove one user by id
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete not implemented!"))
}
