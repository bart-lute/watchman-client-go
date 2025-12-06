package watchman_client_go

import (
	"fmt"
	"net/http"

	"github.com/bart-lute/watchman-client-go/models"
)

var noteFormFields = []string{
	"computer_id",
	"body",
	"include_in_email",
}

/*
ListNotes Get a list of all notes
See: https://api.watchmanmonitoring.com/#list_notes
*/
func (c *Client) ListNotes() (*[]models.Note, error) {
	var notes []models.Note
	err := c.getList("notes", &notes)
	if err != nil {
		return nil, err
	}
	return &notes, nil
}

/*
GetNote Get a single note by its UID
See: https://api.watchmanmonitoring.com/#get_note
*/
func (c *Client) GetNote(uid string) (*models.Note, error) {
	var note models.Note
	err := c.getItem(fmt.Sprintf("notes/%s", uid), &note)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

// CreateNote Create a new note
// See: https://api.watchmanmonitoring.com/#create_note
func (c *Client) CreateNote(noteDataMap *map[string]string) (*models.Note, error) {
	var note models.Note
	err := c.createOrUpdateItem("notes/", http.MethodPost, noteDataMap, nil, &noteFormFields, "note", &note)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

// UpdateNote Update an existing note
// See: https://api.watchmanmonitoring.com/#update_note
func (c *Client) UpdateNote(uid string, noteDataMap *map[string]string) (*models.Note, error) {
	var note models.Note
	err := c.createOrUpdateItem(fmt.Sprintf("notes/%s", uid), http.MethodPut, noteDataMap, nil, &noteFormFields, "note", &note)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

// DeleteNote Delete an existing note
// See: https://api.watchmanmonitoring.com/#delete_note
func (c *Client) DeleteNote(uid string) error {
	return c.deleteItem(fmt.Sprintf("notes/%s", uid))
}
