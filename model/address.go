package model

type Address struct {
	Identifier string `json:"identifier,omitempty"`
	Name       string `json:"name"`
	Holder     *User  `json:"user,omitempty"`
	CEP        string `json:"cep,omitempty"`
	Number     int    `json:"number,omitempty"`
	Complement string `json:"complement,omitempty"`
}

type RawAddress struct {
	Identifier string `json:"identifier,omitempty" db:"identifier"`
	Name       string `json:"name" db:"name"`
	Holder     string `json:"user,omitempty" db:"holder"`
	CEP        string `json:"cep,omitempty" db:"cep"`
	Number     int    `json:"number,omitempty" db:"number"`
	Complement string `json:"complement,omitempty" db:"complement"`
}
