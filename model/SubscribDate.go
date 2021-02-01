package model

type SubscribeDate struct {
	Status int     `json:"status"`
	Dates  []Dates `json:"list"`
}
type Dates struct {
	Date   string `json:"date"`
	Enable bool   `json:"enable"`
}
