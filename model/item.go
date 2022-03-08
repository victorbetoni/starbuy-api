package model

type Item struct {
	Identifier  string  `db:"identifier" json:"identifier,omitempty"`
	Title       string  `db:"title" json:"title,omitempty"`
	Seller      User    `db:"seller" json:"seller,omitempty"`
	Price       float64 `db:"price" json:"price,omitempty"`
	Stock       int     `db:"stock" json:"stock,omitempty"`
	Category    int     `db:"category" json:"category,omitempty"`
	Description string  `db:"description" json:"description,omitempty"`
}

type ItemWithAssets struct {
	Item   Item     `json:"item"`
	Assets []string `json:"assets"`
}

type DatabaseItem struct {
	Identifier  string  `db:"identifier" json:"identifier,omitempty"`
	Title       string  `db:"title" json:"title,omitempty"`
	Seller      string  `db:"seller" json:"seller,omitempty"`
	Price       float64 `db:"price" json:"price,omitempty"`
	Stock       int     `db:"stock" json:"stock,omitempty"`
	Category    int     `db:"category" json:"category,omitempty"`
	Description string  `db:"description" json:"description,omitempty"`
}
