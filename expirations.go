package watchman_client_go

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/bart-lute/watchman-client-go/models"
)

var expirationFormFieldsCreate = []string{
	"license_key",
	"expires_at",
	"expiration_manufacturer_id",
	"expiration_product_id",
	"group_id",
	"computer_id",
	"renewable",
}

var expirationFormFieldsUpdate = []string{
	"notes",
}

/*
ListExpirations Get a list of all expirations
See: https://api.watchmanmonitoring.com/#list_expirations
*/
func (c *Client) ListExpirations() (*[]models.Expiration, error) {
	var expirations []models.Expiration
	err := c.getList("expirations", &expirations)
	if err != nil {
		return nil, err
	}

	// Sort by computer name
	sort.Slice(expirations, func(i, j int) bool {
		return strings.ToLower(expirations[i].Computer) < strings.ToLower(expirations[j].Computer)
	})

	return &expirations, nil
}

/*
GetExpiration Get a single expiration by its UID
See: https://api.watchmanmonitoring.com/#get_expiration
*/
func (c *Client) GetExpiration(uid string) (*models.Expiration, error) {

	var expiration models.Expiration
	err := c.getItem(fmt.Sprintf("expirations/%s", uid), &expiration)
	if err != nil {
		return nil, err
	}
	return &expiration, nil

}

func (c *Client) CreateExpiration(expirationsDataMap *map[string]string) (*models.Expiration, error) {
	var expiration models.Expiration
	err := c.createOrUpdateItem("expirations/", http.MethodPost, expirationsDataMap, nil, &expirationFormFieldsCreate, "expiration", &expiration)
	if err != nil {
		return nil, err
	}
	return &expiration, nil
}

func (c *Client) UpdateExpiration(uid string, expirationsDataMap *map[string]string) (*models.Expiration, error) {
	var expiration models.Expiration
	err := c.createOrUpdateItem(fmt.Sprintf("expirations/%s", uid), http.MethodPut, expirationsDataMap, nil, &expirationFormFieldsUpdate, "expiration", &expiration)
	if err != nil {
		return nil, err
	}
	return &expiration, nil
}

func (c *Client) DeleteExpiration(uid string) error {
	return c.deleteItem(fmt.Sprintf("expirations/%s", uid))
}
