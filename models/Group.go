package models

type Group struct {
    UID                   string `json:"uid"`
    Slug                  string `json:"slug"`
    Name                  string `json:"name"`
    Description           string `json:"description"`
    ReferenceEmail        string `json:"reference_email"`
    Hidden                bool   `json:"hidden"`
    ShowContactMenu       bool   `json:"show_contact_menu"`
    VisibleComputerCount  int    `json:"visible_computer_count"`
    HiddenComputerCount   int    `json:"hidden_computer_count"`
    MacInstallerReady     bool   `json:"mac_installer_ready"`
    MacInstallerURL       string `json:"mac_installer_url"`
    WindowsInstallerReady bool   `json:"windows_installer_ready"`
    WindowsInstallerURL   string `json:"windows_installer_url"`
    CreatedAt             int    `json:"created_at"`
}
