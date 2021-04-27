package repositories

import (
	"backend/src/database"
	"backend/src/models"
	"fmt"
)

type users struct{}

// That factory creates a users repository
func UsersRepository() *users {
	return &users{}
}

func (u users) Create(user models.User) (uint64, error) {
	var ID int

	db, err := database.Connect()

	if err != nil {
		return 0, err
	}

	err = db.QueryRow("insert into tb_users (ds_name, ds_nick, ds_email, ds_password) values($1, $2, $3, $4) RETURNING id", user.Name, user.Nick, user.Email, user.Password).Scan(&ID)

	if err != nil {
		return 0, err
	}

	defer db.Close()

	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

func (u users) Show(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	db, err := database.Connect()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT id, ds_name, ds_nick, ds_email FROM tb_users WHERE LOWER(ds_name) LIKE $1 OR ds_nick LIKE $2", nameOrNick, nameOrNick)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	defer db.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u users) Find(ID uint64) (models.User, error) {
	db, err := database.Connect()
	var user models.User

	if err != nil {
		return models.User{}, err
	}

	err = db.QueryRow("SELECT id, ds_name, ds_nick, ds_email, dt_created_at FROM tb_users WHERE id=$1", ID).Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)

	if err != nil {
		return models.User{}, err
	}

	defer db.Close()

	return user, nil
}
