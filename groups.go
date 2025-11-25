package watchman_client_go

import (
    "fmt"
    "sort"
    "strings"

    "github.com/bart-lute/watchman-client-go/models"
)

/*
ListGroups Get a list of all groups
See: https://api.watchmanmonitoring.com/#list_groups
*/
func (c *Client) ListGroups() *[]models.Group {
    var groups []models.Group
    c.getList("groups", &groups)

    sort.Slice(groups, func(i, j int) bool {
        return strings.ToLower(groups[i].Name) < strings.ToLower(groups[j].Name)
    })
    return &groups
}

/*
GetGroup Get a single group by its UID
See: https://api.watchmanmonitoring.com/#get_group
*/
func (c *Client) GetGroup(uid string) *models.Group {
    var group models.Group
    c.getItem(fmt.Sprintf("groups/%s", uid), &group)
    return &group
}
