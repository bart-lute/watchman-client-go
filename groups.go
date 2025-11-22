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
ListGroups Get a list of all groups
See: https://api.watchmanmonitoring.com/#list_groups
*/
func (c *Client) ListGroups() *[]models.Group {
    var groups []models.Group
    var groupList []models.Group

    page := 1
    for {
        response := c.doRequest("GET", fmt.Sprintf("groups?page=%d", page), nil)
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

        if err := json.Unmarshal(body, &groups); err != nil {
            log.Fatal(err)
        }

        groupList = append(groupList, groups...)

        if page >= totalPages {
            break
        }
        page++

    }

    sort.Slice(groupList, func(i, j int) bool { return strings.ToLower(groupList[i].Name) < strings.ToLower(groupList[j].Name) })
    return &groupList
}

func (c *Client) GetGroup(uid string) *models.Group {
    response := c.doRequest("GET", fmt.Sprintf("/groups/%s", uid), nil)

    // Make sure to close the Body
    defer func(Body io.ReadCloser) {
        err := Body.Close()
        if err != nil {
            log.Fatal(err)
        }
    }(response.Body)

    var group models.Group
    body, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    if err := json.Unmarshal(body, &group); err != nil {
        log.Fatal(err)
    }

    return &group
}
