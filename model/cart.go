package model

type RawCartItem struct {
	Holder   string `db:"holder" json:"holder"`
	Quantity int    `db:"quantity" json:"quantity"`
	Item     string `db:"product" json:"item"`
}

type CartItem struct {
	Holder   *User           `json:"holder,omitempty"`
	Quantity int             `json:"quantity"`
	Item     *ItemWithAssets `json:"item"`
}
