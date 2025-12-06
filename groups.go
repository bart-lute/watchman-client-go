// Package watchman_client_go is a client library for the
// Watchman Monitoring API.
// See https://api.watchmanmonitoring.com/
//
// It tries to follow the API as closely as possible.
package watchman_client_go

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/bart-lute/watchman-client-go/models"
)

var groupFormFields = []string{
	"name",
	"description",
	"show_contact_menu",
	"build_installers",
	"reference_email",
}

// GroupData is used for creating or updating a group.
// BuildInstallers default to false and are only used for creating a group.
// ReferenceEmail is optional and only used for updating a group.
type GroupData struct {
	Name            string
	Description     string
	ShowContactMenu bool   `default:"false"`
	BuildInstallers bool   `default:"false"`
	ReferenceEmail  string `default:""`
}

/*
ListGroups Get a list of all groups
See: https://api.watchmanmonitoring.com/#list_groups
*/
func (c *Client) ListGroups() (*[]models.Group, error) {
	var groups []models.Group
	err := c.getList("groups", &groups)
	if err != nil {
		return nil, err
	}

	sort.Slice(groups, func(i, j int) bool {
		return strings.ToLower(groups[i].Name) < strings.ToLower(groups[j].Name)
	})
	return &groups, nil
}

/*
GetGroup Get a single group by its UID
See: https://api.watchmanmonitoring.com/#get_group
*/
func (c *Client) GetGroup(uid string) (*models.Group, error) {
	var group models.Group
	err := c.getItem(fmt.Sprintf("groups/%s", uid), &group)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

/*
FindOrCreateGroup Find or create a group
*/
func (c *Client) FindOrCreateGroup(groupDataMap *map[string]string) (*models.Group, error) {
	var group models.Group
	err := c.createOrUpdateItem("groups/find_or_create", http.MethodPost, groupDataMap, nil, &groupFormFields, "group", &group)
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (c *Client) UpdateGroup(uid string, groupDataMap *map[string]string) (*models.Group, error) {
	var group models.Group
	err := c.createOrUpdateItem(fmt.Sprintf("groups/%s", uid), http.MethodPut, groupDataMap, nil, &groupFormFields, "group", &group)
	if err != nil {
		return nil, err
	}
	return &group, nil

}
