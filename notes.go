package watchman_client_go

import (
	"fmt"

	"github.com/bart-lute/watchman-client-go/models"
)

/*
ListNotes Get a list of all notes
See: https://api.watchmanmonitoring.com/#list_notes
*/
func (c *Client) ListNotes() *[]models.Note {
	var notes []models.Note
	c.getList("notes", &notes)
	return &notes
}

/*
GetNote Get a single note by its UID
See: https://api.watchmanmonitoring.com/#get_note
*/
func (c *Client) GetNote(uid string) *models.Note {
	var note models.Note
	c.getItem(fmt.Sprintf("notes/%s", uid), &note)
	return &note
}
