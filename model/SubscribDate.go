package model

type SubscribeDate struct {
	Status int     `json:"status"`
	Dates  []Dates `json:"list"`
}
type Dates struct {
	Date   string `json:"date"`
	Enable bool   `json:"enable"`
}

type SubscribeDateDetail struct {
	Date        string
	DateDetails []DateDetail `json:"list"`
}
type DateDetail struct {
	Date         string
	CustomerName string `json:"customer"`
	CustomerId   int    `json:"customerid"`
	StartTime    string `json:"StartTime"`
	EndTime      string `json:"EndTime"`
	Mxid         string `json:"mxid"`
	Qty          int    `json:"qty"`
}
