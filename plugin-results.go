package watchman_client_go

import (
	"fmt"
	"net/http"

	"github.com/bart-lute/watchman-client-go/models"
)

const (
	MuteTypeNone    = "none" // Passing none will clear any existing mutes
	MuteTypeIgnored = "ignored"
	MuteTypeTimed   = "timed"
	MuteTypeTicket  = "ticket"
)

var pluginResultFormFields = []string{
	"mute_type",
	"mute_until",
}

/*
ListPluginResults Get a single computer's plugin results
See: https://api.watchmanmonitoring.com/#plugin_result
*/
func (c *Client) ListPluginResults(watchmanId string) (*models.Computer, error) {
	var pluginResult models.Computer
	err := c.getItem(fmt.Sprintf("computers/%s?expand[]=plugin_results", watchmanId), &pluginResult)
	if err != nil {
		return nil, err
	}
	return &pluginResult, nil
}

// MutePluginResult Mute a plugin result
// See: https://api.watchmanmonitoring.com/#mute_plugin_result
func (c *Client) MutePluginResult(uid string, pluginResultsDataMap *map[string]string) (*models.PluginResult, error) {
	var pluginResult models.PluginResult
	err := c.createOrUpdateItem(fmt.Sprintf("plugin_results/%s", uid), http.MethodPut, pluginResultsDataMap, nil, &pluginResultFormFields, "plugin_result", &pluginResult)
	if err != nil {
		return nil, err
	}
	return &pluginResult, nil
}
