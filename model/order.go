package model

type Order struct {
	Identifier string         `json:"identifier"`
	Customer   User           `json:"customer"`
	Seller     User           `json:"seller"`
	Item       ItemWithAssets `json:"item"`
	Price      float64        `json:"price"`
	Quantity   int            `json:"quantity"`
}

type RawOrder struct {
	Identifier string  `json:"identifier" db:"identifier"`
	Customer   string  `json:"customer" db:"holder"`
	Seller     string  `json:"seller" db:"seller"`
	Item       string  `json:"item" db:"product"`
	Price      float64 `json:"price" db:"price"`
	Quantity   int     `json:"quantity" db:"quantity"`
}

type PurchaseUpdate struct {
}
