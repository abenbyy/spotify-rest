package model


type Artist struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Image  string   `json:"image"`
	Albums []*Album `json:"albums"`
}