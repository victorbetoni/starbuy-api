package model

type RawReview struct {
	User    string `db:"user" json:"username"`
	Item    string `db:"product" json:"item"`
	Message string `db:"msg" json:"message"`
	Rate    int    `db:"rate" json:"rate"`
}

type Review struct {
	User    User           `json:"user"`
	Item    ItemWithAssets `json:"item"`
	Message string         `json:"message"`
	Rate    int            `json:"rate"`
}
