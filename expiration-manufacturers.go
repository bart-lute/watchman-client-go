package watchman_client_go

import (
	"fmt"
	"net/http"

	"github.com/bart-lute/watchman-client-go/models"
)

var expirationManufacturerFormFields = []string{
	"name",
}

/*
ListExpirationManufacturers Get a list of all expiration manufacturers
See: https://api.watchmanmonitoring.com/#list_expiration_manufacturers
*/
func (c *Client) ListExpirationManufacturers() (*[]models.ExpirationManufacturer, error) {
	var expirationManufacturers []models.ExpirationManufacturer
	err := c.getList("expiration_manufacturers", &expirationManufacturers)
	if err != nil {
		return nil, err
	}
	return &expirationManufacturers, nil
}

/*
GetExpirationManufacturer Get a single expiration manufacturer by its UID
See: https://api.watchmanmonitoring.com/#get_expiration_manufacturer
*/
func (c *Client) GetExpirationManufacturer(uid string) (*models.ExpirationManufacturer, error) {
	var expirationManufacturer models.ExpirationManufacturer
	err := c.getItem(fmt.Sprintf("expiration_manufacturers/%s", uid), &expirationManufacturer)
	if err != nil {
		return nil, err
	}
	return &expirationManufacturer, nil
}

// CreateExpirationManufacturer Create a new expiration manufacturer
// See: https://api.watchmanmonitoring.com/#create_expiration_manufacturer
func (c *Client) CreateExpirationManufacturer(expirationManufacturersDataMap *map[string]string) (*models.ExpirationManufacturer, error) {
	var expirationManufacturer models.ExpirationManufacturer
	err := c.createOrUpdateItem("expiration_manufacturers/", http.MethodPost, expirationManufacturersDataMap, nil, &expirationManufacturerFormFields, "expiration_manufacturer", &expirationManufacturer)
	if err != nil {
		return nil, err
	}
	return &expirationManufacturer, nil
}

// UpdateExpirationManufacturer Update an existing expiration manufacturer
// See: https://api.watchmanmonitoring.com/#update_expiration_manufacturer
func (c *Client) UpdateExpirationManufacturer(uid string, expirationManufacturersDataMap *map[string]string) (*models.ExpirationManufacturer, error) {
	var expirationManufacturer models.ExpirationManufacturer
	err := c.createOrUpdateItem(fmt.Sprintf("expiration_manufacturers/%s", uid), http.MethodPut, expirationManufacturersDataMap, nil, &expirationManufacturerFormFields, "expiration_manufacturer", &expirationManufacturer)
	if err != nil {
		return nil, err
	}
	return &expirationManufacturer, nil
}

// DeleteExpirationManufacturer Delete an existing expiration manufacturer
// See: https://api.watchmanmonitoring.com/#delete_expiration_manufacturer
func (c *Client) DeleteExpirationManufacturer(uid string) error {
	return c.deleteItem(fmt.Sprintf("expiration_manufacturers/%s", uid))
}
