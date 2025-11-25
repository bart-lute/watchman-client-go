package watchman_client_go

import (
    "fmt"
    "sort"
    "strings"

    "github.com/bart-lute/watchman-client-go/models"
)

/*
ListComputers Get a list of all computers
See: https://api.watchmanmonitoring.com/#list_computers
*/
func (c *Client) ListComputers() *[]models.Computer {
    var computers []models.Computer
    c.getList("computers", &computers)

    sort.Slice(computers, func(i, j int) bool {
        return strings.ToLower(computers[i].ComputerName) < strings.ToLower(computers[j].ComputerName)
    })
    return &computers
}

/*
GetComputer Get a single computer by its Watchman ID
See: https://api.watchmanmonitoring.com/#get_computer
*/
func (c *Client) GetComputer(watchmanId string) *models.Computer {
    var computer models.Computer
    c.getItem(fmt.Sprintf("computers/%s", watchmanId), &computer)
    return &computer
}
