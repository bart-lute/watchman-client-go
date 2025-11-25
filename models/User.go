package models

type User struct {
    UID                 string   `json:"uid"`
    Email               string   `json:"email"`
    Firstname           string   `json:"firstname"`
    Lastname            string   `json:"lastname"`
    TfaStatus           bool     `json:"tfa_status"`
    Role                string   `json:"role"`
    CanEditContactMenu  bool     `json:"can_edit_contact_menu"`
    CanAccessBilling    bool     `json:"can_access_billing"`
    LastSubdomain       string   `json:"last_subdomain"`
    LastSignin          int      `json:"last_signin"`
    CreatedAt           int      `json:"created_at"`
    RestrictedGroupUids []string `json:"restricted_group_uids"`
}
