package model

type CustomerList struct {
	Customers []Customer `json:"list"`
	//Status    int        `json:"status"`
}
type Customer struct {
	Id    int    `json:"id"`
	Cname string `json:"cname"`
	Addr  string `json:"addr"`
	//SmallPic string   `json:"-"`
	//Lat      float64   `json:"-"`
	//Lng      float64   `json:"-"`
	Tel      string `json:"tel"`
	Addr2    string `json:"addr2"`
	Province int    `json:"province"`
	City     int    `json:"city"`
	//County   int      `json:"county"`
	//Sort     int      `json:"-"`
	//Distance float64  `json:"-"`
	//Tags     []string `json:"-"`
}
