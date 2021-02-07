package model

type VerifyPicModel struct {
	Dragon string `json:"dragon"`
	Tiger  string `json:"tiger"`
}

type VerifyResultModel struct {
	Guid   string `json:"guid" default:""`
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
