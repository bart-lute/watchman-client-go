package models

type PluginResult struct {
	UID             string `json:"uid"`
	Name            string `json:"name"`
	Slug            string `json:"slug"`
	Status          string `json:"status"`
	Details         string `json:"details"`
	LastRun         int    `json:"last_run"`
	Muted           bool   `json:"muted"`
	MuteType        string `json:"mute_type"`
	MuteDescription string `json:"mute_description"`
	MutedUntil      int    `json:"muted_until"`
}
