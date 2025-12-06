package watchman_client_go

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/bart-lute/watchman-client-go/models"
)

// computerFormFields contains a list of field that can be used for updating a computer.
// If a field does not appear in this list, it is not supported by the API.
var computerFormFields = []string{
	"asset_id",
	"custom_name",
	"description",
	"reference_email",
	"product_description",
	"estimated_manufacture_date",
	"beacon_reporting",
	"beacon_missing_threshold_in_minutes",
	"create_ticket_on_beacon_found",
}

type ComputerData struct {
	AssetId                         string
	CustomName                      string
	Description                     string
	ReferenceEmail                  string `default:""`
	ProductDescription              string
	EstimatedManufactureDate        string
	BeaconReporting                 bool `default:"false"`
	BeaconMissingThresholdInMinutes int  `default:"5"`
	CreateTicketOnBeaconFound       bool `default:"true"`
}

/*
ListComputers Get a list of all computers
See: https://api.watchmanmonitoring.com/#list_computers
*/
func (c *Client) ListComputers() (*[]models.Computer, error) {
	var computers []models.Computer
	err := c.getList("computers", &computers)
	if err != nil {
		return nil, err
	}

	sort.Slice(computers, func(i, j int) bool {
		return strings.ToLower(computers[i].ComputerName) < strings.ToLower(computers[j].ComputerName)
	})
	return &computers, nil
}

/*
GetComputer Get a single computer by its Watchman ID
See: https://api.watchmanmonitoring.com/#get_computer
*/
func (c *Client) GetComputer(watchmanId string) (*models.Computer, error) {
	var computer models.Computer
	err := c.getItem(fmt.Sprintf("computers/%s", watchmanId), &computer)
	if err != nil {
		return nil, err
	}
	return &computer, nil
}

/*
UpdateComputer Update a computer with a subset of fields
See: https://api.watchmanmonitoring.com/#update_computer
Valid computerDataMap keys are:
asset_id
custom_name
description
reference_email
product_description
estimated_manufacture_date
beacon_reporting
beacon_missing_threshold_in_minutes
create_ticket_on_beacon_found
*/
func (c *Client) UpdateComputer(watchmanID string, computerDataMap *map[string]string) (*models.Computer, error) {
	var computer models.Computer
	err := c.createOrUpdateItem(fmt.Sprintf("computers/%s", watchmanID), http.MethodPut, computerDataMap, nil, &computerFormFields, "computer", &computer)
	if err != nil {
		return nil, err
	}
	return &computer, nil
}

/*
UpdateComputerAll Update a computer with all available fields
It is essentially replacing all fields with the values provided in the request body.
Not sure if this is useful, but it's here if you need it.'
*/
func (c *Client) UpdateComputerAll(watchmanID string, computerData *ComputerData) (*models.Computer, error) {
	requestBody := map[string]string{
		"asset_id":                            computerData.AssetId,
		"custom_name":                         computerData.CustomName,
		"description":                         computerData.Description,
		"reference_email":                     computerData.ReferenceEmail,
		"product_description":                 computerData.ProductDescription,
		"estimated_manufacture_date":          computerData.EstimatedManufactureDate,
		"beacon_reporting":                    strconv.FormatBool(computerData.BeaconReporting),
		"beacon_missing_threshold_in_minutes": strconv.Itoa(computerData.BeaconMissingThresholdInMinutes),
		"create_ticket_on_beacon_found":       strconv.FormatBool(computerData.CreateTicketOnBeaconFound),
	}

	return c.UpdateComputer(watchmanID, &requestBody)
}
