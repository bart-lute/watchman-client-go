package models

type ExpirationManufacturer struct {
	UID       string `json:"uid"`
	Slug      string `json:"slug"`
	Name      string `json:"name"`
	CreatedAt int    `json:"created_at"`
}
