package watchman_client_go

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/bart-lute/watchman-client-go/models"
)

var userFormFields = []string{
	"firstname",
	"lastname",
	"email",
	"role",
	"can_edit_contact_menu",
	"can_access_billing",
}

/*
ListUsers Get a list of all users
See: https://api.watchmanmonitoring.com/#list_users
*/
func (c *Client) ListUsers() (*[]models.User, error) {

	var users []models.User
	err := c.getList("users", &users)
	if err != nil {
		return nil, err
	}

	// Sort by lastname, firstname
	sort.Slice(users, func(i, j int) bool {
		return strings.ToLower(fmt.Sprintf("%s, %s", users[i].Lastname, users[i].Firstname)) < strings.ToLower(fmt.Sprintf("%s, %s", users[j].Lastname, users[j].Firstname))
	})

	return &users, nil
}

/*
GetUser Get a single user by its UID
See: https://api.watchmanmonitoring.com/#get_user
*/
func (c *Client) GetUser(uid string) (*models.User, error) {
	var user models.User
	err := c.getItem(fmt.Sprintf("users/%s", uid), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser Find or create a user
// See: https://api.watchmanmonitoring.com/#create_user
func (c *Client) CreateUser(userDataMap *map[string]string, restrictedGroupUIDs *[]string) (*models.User, error) {
	var user models.User
	endpoint := "users"
	additionalFormData := make(map[string][]string)
	if restrictedGroupUIDs != nil {
		additionalFormData["user[restricted_group_uids][]"] = *restrictedGroupUIDs
	}
	err := c.createOrUpdateItem(endpoint, http.MethodPost, userDataMap, &additionalFormData, &userFormFields, "user", &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
