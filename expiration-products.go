package watchman_client_go

import (
	"fmt"

	"github.com/bart-lute/watchman-client-go/models"
)

/*
ListExpirationProducts Get a list of all expiration products
See: https://api.watchmanmonitoring.com/#list_expiration_products
*/
func (c *Client) ListExpirationProducts() *[]models.ExpirationProduct {
	var expirationProducts []models.ExpirationProduct
	c.getList("expiration_products", &expirationProducts)

	return &expirationProducts
}

/*
GetExpirationProduct Get a single expiration product by its UID
See: https://api.watchmanmonitoring.com/#get_expiration_product
*/
func (c *Client) GetExpirationProduct(uid string) *models.ExpirationProduct {
	var expirationProduct models.ExpirationProduct
	c.getItem(fmt.Sprintf("expiration_products/%s", uid), &expirationProduct)
	return &expirationProduct
}
