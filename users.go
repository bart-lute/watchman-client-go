package watchman_client_go

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bart-lute/watchman-client-go/models"
)

/*
ListUsers Get a list of all users
See: https://api.watchmanmonitoring.com/#list_users
*/
func (c *Client) ListUsers() *[]models.User {

	var users []models.User
	c.getList("users", &users)

	sort.Slice(users, func(i, j int) bool {
		return strings.ToLower(fmt.Sprintf("%s, %s", users[i].Lastname, users[i].Firstname)) < strings.ToLower(fmt.Sprintf("%s, %s", users[j].Lastname, users[j].Firstname))
	})
	return &users
}

/*
GetUser Get a single user by its UID
See: https://api.watchmanmonitoring.com/#get_user
*/
func (c *Client) GetUser(uid string) *models.User {
	var user models.User
	c.getItem(fmt.Sprintf("users/%s", uid), &user)
	return &user
}
