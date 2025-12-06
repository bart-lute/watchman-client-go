package watchman_client_go

import (
	"fmt"
	"net/http"

	"github.com/bart-lute/watchman-client-go/models"
)

var expirationProductFormFields = []string{
	"name",
}

/*
ListExpirationProducts Get a list of all expiration products
See: https://api.watchmanmonitoring.com/#list_expiration_products
*/
func (c *Client) ListExpirationProducts() (*[]models.ExpirationProduct, error) {
	var expirationProducts []models.ExpirationProduct
	err := c.getList("expiration_products", &expirationProducts)
	if err != nil {
		return nil, err
	}

	return &expirationProducts, nil
}

/*
GetExpirationProduct Get a single expiration product by its UID
See: https://api.watchmanmonitoring.com/#get_expiration_product
*/
func (c *Client) GetExpirationProduct(uid string) (*models.ExpirationProduct, error) {
	var expirationProduct models.ExpirationProduct
	err := c.getItem(fmt.Sprintf("expiration_products/%s", uid), &expirationProduct)
	if err != nil {
		return nil, err
	}
	return &expirationProduct, nil
}

// CreateExpirationProduct Create a new expiration product
// See: https://api.watchmanmonitoring.com/#create_expiration_product
func (c *Client) CreateExpirationProduct(expirationProductsDataMap *map[string]string) (*models.ExpirationProduct, error) {
	var expirationProduct models.ExpirationProduct
	err := c.createOrUpdateItem("expiration_products/", http.MethodPost, expirationProductsDataMap, nil, &expirationProductFormFields, "expiration_product", &expirationProduct)
	if err != nil {
		return nil, err
	}
	return &expirationProduct, nil
}

// UpdateExpirationProduct Update an existing expiration product
// See: https://api.watchmanmonitoring.com/#update_expiration_product
func (c *Client) UpdateExpirationProduct(uid string, expirationProductsDataMap *map[string]string) (*models.ExpirationProduct, error) {
	var expirationProduct models.ExpirationProduct
	err := c.createOrUpdateItem(fmt.Sprintf("expiration_products/%s", uid), http.MethodPut, expirationProductsDataMap, nil, &expirationProductFormFields, "expiration_product", &expirationProduct)
	if err != nil {
		return nil, err
	}
	return &expirationProduct, nil
}

// DeleteExpirationProduct Delete an existing expiration product
// See: https://api.watchmanmonitoring.com/#delete_expiration_product
func (c *Client) DeleteExpirationProduct(uid string) error {
	return c.deleteItem(fmt.Sprintf("expiration_products/%s", uid))
}
