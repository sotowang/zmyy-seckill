package model

type City struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Children []City `json:"children"`
}
