package watchman_client_go

import (
	"fmt"

	"github.com/bart-lute/watchman-client-go/models"
)

/*
PluginResults Get a single computer's plugin results
See: https://api.watchmanmonitoring.com/#plugin_result
*/
func (c *Client) PluginResults(watchmanId string) *models.Computer {
	var pluginResult models.Computer
	c.getItem(fmt.Sprintf("computers/%s?expand[]=plugin_results", watchmanId), &pluginResult)
	return &pluginResult
}
