package watchman_client_go

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "sort"
    "strconv"
    "strings"

    "github.com/bart-lute/watchman-client-go/models"
)

/*
ListComputers Get a list of all computers
See: https://api.watchmanmonitoring.com/#list_computers
*/
func (c *Client) ListComputers() *[]models.Computer {
    var computers []models.Computer
    var computerList []models.Computer

    page := 1
    for {
        response := c.doRequest("GET", fmt.Sprintf("computers?page=%d", page), nil)
        totalPages, err := strconv.Atoi(response.Header.Get("x-total-pages"))
        if err != nil {
            log.Fatal(err)
        }

        body, err := io.ReadAll(response.Body)
        if err != nil {
            log.Fatal(err)
        }

        // Close the Body
        err = response.Body.Close()
        if err != nil {
            log.Fatal(err)
        }

        if err := json.Unmarshal(body, &computers); err != nil {
            log.Fatal(err)
        }

        computerList = append(computerList, computers...)

        if page >= totalPages {
            break
        }
        page++
    }
    sort.Slice(computerList, func(i, j int) bool { return strings.ToLower(computerList[i].ComputerName) < strings.ToLower(computerList[j].ComputerName) })
    return &computerList
}

/*
GetComputer Get a single computer by its Watchman ID
See: https://api.watchmanmonitoring.com/#get_computer
*/
func (c *Client) GetComputer(watchmanId string) *models.Computer {
    response := c.doRequest("GET", fmt.Sprintf("/computers/%s", watchmanId), nil)

    // Make sure to close the Body
    defer func(Body io.ReadCloser) {
        err := Body.Close()
        if err != nil {
            log.Fatal(err)
        }
    }(response.Body)

    var computer models.Computer
    body, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    if err := json.Unmarshal(body, &computer); err != nil {
        log.Fatal(err)
    }

    return &computer
}
