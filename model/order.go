package model

type Order struct {
	Identifier string         `json:"identifier"`
	Customer   User           `json:"customer"`
	Seller     User           `json:"seller"`
	Item       ItemWithAssets `json:"item"`
	Price      float64        `json:"price"`
	Quantity   int            `json:"quantity"`
	Status     int            `json:"status"`
	SendTo     Address        `json:"send_to"`
	Date       string         `json:"date"`
}

type OrderWithItem struct {
	Order Order          `json:"order"`
	Item  ItemWithAssets `json:"item"`
}

type RawOrder struct {
	Identifier string  `json:"identifier" db:"identifier"`
	Customer   string  `json:"customer" db:"holder"`
	Seller     string  `json:"seller" db:"seller"`
	Item       string  `json:"item" db:"product"`
	Price      float64 `json:"price" db:"price"`
	Quantity   int     `json:"quantity" db:"amount"`
	Status     int     `json:"status" db:"status"`
	SendTo     string  `json:"send_to" db:"send_to"`
	Date       string  `json:"date" db:"date"`
}
