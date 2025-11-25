package watchman_client_go

import (
	"fmt"

	"github.com/bart-lute/watchman-client-go/models"
)

/*
ListExpirationManufacturers Get a list of all expiration manufacturers
See: https://api.watchmanmonitoring.com/#list_expiration_manufacturers
*/
func (c *Client) ListExpirationManufacturers() *[]models.ExpirationManufacturer {
	var expirationManufacturers []models.ExpirationManufacturer
	c.getList("expiration_manufacturers", &expirationManufacturers)
	return &expirationManufacturers
}

/*
GetExpirationManufacturer Get a single expiration manufacturer by its UID
See: https://api.watchmanmonitoring.com/#get_expiration_manufacturer
*/
func (c *Client) GetExpirationManufacturer(uid string) *models.ExpirationManufacturer {
	var expirationManufacturer models.ExpirationManufacturer
	c.getItem(fmt.Sprintf("expiration_manufacturers/%s", uid), &expirationManufacturer)
	return &expirationManufacturer
}
