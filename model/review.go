package model

type RawReview struct {
	User    string `db:"user" json:"user"`
	Item    string `db:"product" json:"item"`
	Message string `db:"msg "json:"message"`
	Rate    int    `db:"rate" json:"rate"`
}
