package model

type Address struct {
	Identifier string `json:"identifier,omitempty"`
	Holder     *User  `json:"user,omitempty"`
	CEP        string `json:"cep,omitempty"`
	Number     int    `json:"number,omitempty"`
	Complement string `json:"complement,omitempty"`
}

type RawAddress struct {
	Identifier string `json:"identifier,omitempty" db:"identifier"`
	Holder     string `json:"user,omitempty" db:"holder"`
	CEP        string `json:"cep,omitempty" db:"cep"`
	Number     int    `json:"number,omitempty" db:"number"`
	Complement string `json:"complement,omitempty" db:"complement"`
}
