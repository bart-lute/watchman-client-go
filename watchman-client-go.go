package watchman_client_go

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"
    "strconv"
    "strings"
    "time"
)

type Client struct {
    baseUrl    string
    apiKey     string
    httpClient *http.Client
}

const timeout = 10 * time.Second

func Init(baseUrl string, apiKey string) *Client {
    return &Client{baseUrl, apiKey, &http.Client{Timeout: timeout}}
}

func (c *Client) doRequest(method string, endPoint string, requestBody any, requestHeaders map[string]string) *http.Response {

    // Whatever the endPoint is, we always want to add the API Key to the query parameters
    u, err := url.Parse(fmt.Sprintf("%s/%s", c.baseUrl, endPoint))
    if err != nil {
        log.Fatal(err)
    }
    q := u.Query()
    q.Add("api_key", c.apiKey)
    u.RawQuery = q.Encode()

    // Convert the Request Body to a Reader Object
    var requestBodyReader io.Reader
    if requestBody != nil {
        rb, err := json.Marshal(requestBody)
        if err != nil {
            log.Fatal(err)
        }
        requestBodyReader = strings.NewReader(string(rb))
    }

    request, err := http.NewRequest(method, u.String(), requestBodyReader)
    if err != nil {
        log.Fatal(err)
    }

    request.Header.Add("Accept", "application/json")
    request.Header.Add("Content-Type", "application/json")
    if requestHeaders != nil {
        for key, value := range requestHeaders {
            request.Header.Add(key, value)
        }
    }

    response, err := c.httpClient.Do(request)
    if err != nil {
        log.Fatal(err)
    }

    if response.StatusCode != http.StatusOK {
        log.Fatal(fmt.Sprintf("API error: %s", response.Status))
    }

    return response
}

func (c *Client) getList(endpoint string, responseBody any) {

    var items []any
    page := 1
    for {
        response := c.doRequest("GET", fmt.Sprintf("%s?page=%d", endpoint, page), nil, nil)
        totalPages, err := strconv.Atoi(response.Header.Get("x-total-pages"))
        if err != nil {
            log.Fatal(err)
        }

        body, err := io.ReadAll(response.Body)
        if err != nil {
            log.Fatal(err)
        }
        err = response.Body.Close()
        if err != nil {
            log.Fatal(err)
        }

        var pageItems []any
        err = json.Unmarshal(body, &pageItems)
        if err != nil {
            log.Fatal(err)
        }
        items = append(items, pageItems...)

        if page >= totalPages {
            break
        }
        page++
    }

    bytes, err := json.Marshal(items)
    if err != nil {
        log.Fatal(err)
    }

    err = json.Unmarshal(bytes, responseBody)
    if err != nil {
        log.Fatal(err)
    }

}

func (c *Client) getItem(endpoint string, responseBody any) {
    response := c.doRequest("GET", endpoint, nil, nil)

    // Make sure to close the Body
    defer func(Body io.ReadCloser) {
        err := Body.Close()
        if err != nil {
            log.Fatal(err)
        }
    }(response.Body)

    body, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }

    if err := json.Unmarshal(body, responseBody); err != nil {
        log.Fatal(err)
    }

}
