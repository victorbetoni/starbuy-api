package model

type RawReview struct {
	User    string `db:"username" json:"user"`
	Item    string `db:"product" json:"item"`
	Message string `db:"message" json:"message"`
	Rate    int    `db:"rate" json:"rate"`
}

type Review struct {
	User    User            `json:"user,omitempty"`
	Item    *ItemWithAssets `json:"item,omitempty"`
	Message string          `json:"message,omitempty"`
	Rate    int             `json:"rate,omitempty"`
}
