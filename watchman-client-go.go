package watchman_client_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Client for interacting with the Watchman API
type Client struct {
	BaseUrl    string
	ApiKey     string
	httpClient http.Client
}

const timeout = 10 * time.Second

func (c *Client) doRequest(method string, endPoint string, urlValues *url.Values, requestHeaders map[string]string) (*http.Response, error) {

	c.httpClient.Timeout = timeout

	// Whatever the endPoint is, we always want to add the API Key to the query parameters
	u, err := url.Parse(fmt.Sprintf("%s/%s", c.BaseUrl, endPoint))
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Add("api_key", c.ApiKey)
	u.RawQuery = q.Encode()

	// Convert the Request Body to a Reader Object
	var requestBodyReader io.Reader
	if urlValues != nil {
		//data := url.Values{}
		//for key, value := range *requestBody {
		//	data.Add(key, value)
		//}
		//requestBodyReader = strings.NewReader(data.Encode())
		requestBodyReader = strings.NewReader(urlValues.Encode())
	}

	request, err := http.NewRequest(method, u.String(), requestBodyReader)
	if err != nil {
		return nil, err
	}

	// Add any additional headers
	if requestHeaders != nil {
		for key, value := range requestHeaders {
			request.Header.Add(key, value)
		}
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	// Check for API errors (< 200 || >= 300)
	if (response.StatusCode < http.StatusOK) || (response.StatusCode >= http.StatusMultipleChoices) {
		return nil, errors.New(fmt.Sprintf("API error: %s", response.Status))
	}

	return response, nil
}

func (c *Client) getList(endpoint string, responseBody any) error {

	var items []any
	page := 1
	for {
		response, err := c.doRequest("GET", fmt.Sprintf("%s?page=%d", endpoint, page), nil, nil)
		if err != nil {
			return err
		}
		totalPages, err := strconv.Atoi(response.Header.Get("x-total-pages"))
		if err != nil {
			return err
		}

		body, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		err = response.Body.Close()
		if err != nil {
			return err
		}

		var pageItems []any
		err = json.Unmarshal(body, &pageItems)
		if err != nil {
			return err
		}
		items = append(items, pageItems...)

		if page >= totalPages {
			break
		}
		page++
	}

	bytes, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, responseBody)
	if err != nil {
		return err
	}

	return nil
}

func getResponseBody(response *http.Response, responseBody any) error {

	// Make sure to close the Body
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, responseBody); err != nil {
		return err
	}

	return nil
}

func (c *Client) getItem(endpoint string, responseBody any) error {
	response, err := c.doRequest(http.MethodGet, endpoint, nil, nil)
	if err != nil {
		return err
	}
	err = getResponseBody(response, responseBody)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) createOrUpdateItem(
	endpoint string,
	method string,
	dataMap *map[string]string, // The data to send in the request body
	additionalFormData *map[string][]string,
	formFields *[]string,
	prefix string,
	responseBody any,
) error {

	urlValues, err := createFormData(formFields, dataMap, prefix)
	if err != nil {
		return err
	}

	// Add any additional form data
	// Note that this data is not validated, so the caller must ensure that it is valid
	if additionalFormData != nil {
		for key, values := range *additionalFormData {
			for _, value := range values {
				urlValues.Add(key, value)
			}
		}
	}

	requestHeaders := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	response, err := c.doRequest(method, endpoint, urlValues, requestHeaders)
	if err != nil {
		return err
	}
	err = getResponseBody(response, responseBody)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) deleteItem(endpoint string) error {
	_, err := c.doRequest(http.MethodDelete, endpoint, nil, nil)
	return err
}

func createFormData(formFields *[]string, dataMap *map[string]string, prefix string) (*url.Values, error) {
	for key := range *dataMap {
		if !slices.Contains(*formFields, key) {
			return nil, fmt.Errorf("'%s' is not a valid field", key)
		}
	}

	// Create a new map with the prefix added to each key
	urlValues := url.Values{}
	for key, value := range *dataMap {
		urlValues.Add(fmt.Sprintf("%s[%s]", prefix, key), value)
		//formattedDataMap[fmt.Sprintf("%s[%s]", prefix, key)] = value
	}

	return &urlValues, nil
}
