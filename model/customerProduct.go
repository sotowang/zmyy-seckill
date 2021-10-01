package model

type RootSource struct {
	Status           int               `json:"status"`
	StartDate        string            `json:"startDate"`
	EndDate          string            `json:"endDate"`
	Memo             string            `json:"-"`
	Tel              string            `json:"tel"`
	Addr             string            `json:"addr"`
	Cname            string            `json:"cname"`
	Lat              float64           `json:"-"`
	Lng              float64           `json:"-"`
	Distance         float64           `json:"distance"`
	CustomerProducts []CustomerProduct `json:"list"`
}
type CustomerProduct struct {
	Id              int       `json:"id"`
	Text            string    `json:"text"`
	Price           string   `json:"price"`
	Descript        string    `json:"descript"`
	Warn            string    `json:"warn"`
	Tags            []string  `json:"tags"`
	QuestionnaireId int       `json:"questionnaireId"`
	Remarks         string    `json:"remarks"`
	NumbersVaccine  []Vaccine `json:"NumbersVaccine"`
	Date            string    `json:"date"`
	BtnLable        string    `json:"BtnLable"`
	Enable          bool      `json:"enable"`
}
type Payment struct {
	BigPic      string `json:"big_pic"`
	IdcardLimit bool   `json:"IdcardLimit"`
	Notice      string `json:"notice"`
}
type Vaccine struct {
	Cname string `json:"cname"`
	Value int    `json:"value"`
}
