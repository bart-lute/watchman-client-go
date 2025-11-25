package models

type Expiration struct {
    UID                    string `json:"uid"`
    LicenseKey             string `json:"license_key"`
    ExpiresAt              int    `json:"expires_at"`
    ExpirationManufacturer string `json:"expiration_manufacturer"`
    ExpirationProduct      string `json:"expiration_product"`
    Computer               string `json:"computer"`
    Group                  string `json:"group"`
    Renewable              bool   `json:"renewable"`
    Notes                  string `json:"notes"`
    CreatedAt              int    `json:"created_at"`
}
