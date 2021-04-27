package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (u *User) Prepare() error {
	if err := u.validate(); err != nil {
		return err
	}

	u.format()
	return nil
}

func (u *User) validate() error {
	var errs []string

	if u.Name == "" {
		errs = append(errs, "name is required")
	}

	if u.Nick == "" {
		errs = append(errs, "nick is required")
	}

	if u.Email == "" {
		errs = append(errs, "email is required")
	}

	if u.Password == "" {
		errs = append(errs, "password is required")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, ","))
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
}
