package watchman_client_go

import (
    "fmt"
    "sort"
    "strings"

    "github.com/bart-lute/watchman-client-go/models"
)

/*
ListExpirations Get a list of all expirations
See: https://api.watchmanmonitoring.com/#list_expirations
*/
func (c *Client) ListExpirations() *[]models.Expiration {
    var expirations []models.Expiration
    c.getList("expirations", &expirations)
    sort.Slice(expirations, func(i, j int) bool {
        return strings.ToLower(expirations[i].Computer) < strings.ToLower(expirations[j].Computer)
    })
    return &expirations
}

/*
GetExpiration Get a single expiration by its UID
See: https://api.watchmanmonitoring.com/#get_expiration
*/
func (c *Client) GetExpiration(uid string) *models.Expiration {

    var expiration models.Expiration
    c.getItem(fmt.Sprintf("expirations/%s", uid), &expiration)
    return &expiration

}
