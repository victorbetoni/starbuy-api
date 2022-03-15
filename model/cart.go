package model

type RawCartItem struct {
	Holder   string `db:"holder", json:"holder"`
	Quantity int    `db:"quantity", json:"quantity"`
	Item     string `db:"item", json:"item"`
}

type CartItem struct {
	Holder   string `json:"holder"`
	Quantity int    `json:"quantity"`
	Item     Item   `json:"item"`
}
