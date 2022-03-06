package model

type User struct {
	Username       string `db:"username" 		json:"username,omitempty"`
	Email          string `db:"email" 			json:"email,omitempty"`
	Name           string `db:"name" 			json:"name,omitempty"`
	Gender         int    `db:"gender" 			json:"gender,omitempty"`
	Registration   string `db:"registration" 	json:"registration,omitempty"`
	Birthdate      string `db:"birthdate" 		json:"birthdate,omitempty"`
	Seller         bool   `db:"seller" 			json:"seller,omitempty"`
	ProfilePicture string `db:"profile_picture" json:"profile_picture,omitempty"`
	City           string `db:"city" 			json:"city,omitempty"`
}
