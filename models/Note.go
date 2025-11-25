package models

type Note struct {
	UID             string `json:"uid"`
	Computer        string `json:"computer"`
	Group           string `json:"group"`
	Body            string `json:"body"`
	CreatedAt       int    `json:"created_at"`
	Private         bool   `json:"private"`
	IncludeInEmail  bool   `json:"include_in_email"`
	IncludeInReport bool   `json:"include_in_report"`
}
