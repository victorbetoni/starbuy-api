package model

import (
	"errors"
	"strings"
)

type User struct {
	Username       string `db:"username" json:"username,omitempty"`
	Email          string `db:"email" json:"email,omitempty"`
	Name           string `db:"name" json:"name,omitempty"`
	Birthdate      string `db:"birthdate" json:"birthdate,omitempty"`
	Seller         bool   `db:"seller" json:"seller"`
	ProfilePicture string `db:"profile_picture" json:"profile_picture,omitempty"`
	City           string `db:"city" json:"city,omitempty"`
	Registration   string `db:"registration" json:"registration,omitempty"`
}

func (user *User) Prepare() error {
	user.format()
	if err := user.validate(); err != nil {
		return err
	}

	return nil
}

func (user *User) validate() error {
	if user.Username == "" {
		return errors.New("Username inválido")
	}
	if user.City == "" {
		return errors.New("Cidade inválida")
	}
	if user.Email == "" {
		return errors.New("Email inválido")
	}
	if user.Name == "" {
		return errors.New("Nome inválido")
	}

	return nil
}

func (user *User) format() {
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)
	user.Name = strings.TrimSpace(user.Name)
	user.City = strings.TrimSpace(user.City)
}
