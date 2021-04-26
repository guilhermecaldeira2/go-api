package repositories

import (
	"backend/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

// That factory creates a users repository
func UsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) Create(user models.User) (uint64, error) {
	err := u.db.QueryRow("insert into tb_users (ds_name, ds_nick, ds_email, ds_password) values($1, $2, $3, $4) RETURNING id", user.Name, user.Nick, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return 0, err
	}

	defer u.db.Close()

	if err != nil {
		return 0, err
	}

	return uint64(user.ID), nil
}
