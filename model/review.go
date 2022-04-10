package model

type RawReview struct {
	Identifier string `db:"identifier" json:"identifier"`
	User       string `db:"user" json:"user"`
	Item       string `db:"product" json:"item"`
	Message    string `db:"msg" json:"message"`
	Rate       int    `db:"rate" json:"rate"`
}

type Review struct {
	Identifier string         `json:"identifier"`
	User       User           `json:"user"`
	Item       ItemWithAssets `json:"item"`
	Message    string         `json:"message"`
	Rate       int            `json:"rate"`
}
