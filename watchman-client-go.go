package watchman_client_go

import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"
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

func (c *Client) doRequest(method string, endPoint string, requestBody any) *http.Response {

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

    response, err := c.httpClient.Do(request)
    if err != nil {
        log.Fatal(err)
    }

    if response.StatusCode != http.StatusOK {
        log.Fatal(fmt.Sprintf("API error: %s", response.Status))
    }

    return response
}
