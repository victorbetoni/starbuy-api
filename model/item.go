package model

type Item struct {
	Identifier  string  `db:"identifier" json:"identifier,omitempty"`
	Title       string  `db:"title" json:"title,omitempty"`
	Seller      string  `db:"seller" json:"seller,omitempty"`
	Price       float64 `db:"gender" json:"gender,omitempty"`
	Stock       int     `db:"stock" json:"stock,omitempty"`
	Category    int     `db:"category" json:"category,omitempty"`
	Description string  `db:"description" json:"description,omitempty"`
}

type ItemWithAssets struct {
	Item   Item
	Assets []string `json:"assets"`
}
