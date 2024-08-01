package models

import "time"

type Admin struct {
	Id       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Event struct {
	Id           uint          `gorm:"primaryKey"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	InitDate     time.Time     `json:"init_date"`
	FinishDate   time.Time     `json:"finish_date"`
	Location     string        `json:"location"`
	TicketPrice  float32       `json:"ticket_price"`
	ImageUrl     string        `json:"image_url"`
	Participants []Participant `json:"participants"`
}

type Participant struct {
	Id           uint   `gorm:"primaryKey"`
	ProfileImage string `json:"profile_image"`
	DocumentFile string `json:"document_file"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Age          uint   `json:"age"`
	Email        string `json:"email"`
	Contact      string `json:"contact"`
}
