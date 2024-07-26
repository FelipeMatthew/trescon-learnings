package model

type Library struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
}

func (Library) TableName() string {
	return "library"
}
