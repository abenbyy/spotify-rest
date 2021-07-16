package model

type Album struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Image  string   `json:"image"`
	Tracks []*Track `json:"tracks"`
}