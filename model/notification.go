package model

type RawNotification struct {
	Identifier string `json:"identifier,omitempty"`
	User       string `json:"target,omitempty"`
	Text       string `json:"text,omitempty"`
	SentIn     string `json:"sent_in,omitempty"`
}

type Notification struct {
	Identifier string `json:"identifier,omitempty"`
	User       *User  `json:"target,omitempty"`
	Text       string `json:"text,omitempty"`
	SentIn     string `json:"sent_in,omitempty"`
}
