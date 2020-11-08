package models

type Permission struct {
	Id     int64  `json:"id"`
	Source string `json:"source"`
	Mode   string `json:"mode"`
}
