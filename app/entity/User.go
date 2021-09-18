package entity

type User struct {
	ID            uint64 `json:"id"`
	Username      string `json:"username"`
	DisplayName   string `json:"display_name"`
	TwitterUserID string `json:"-"`
	IconURL       string `json:"icon_url"`
}
