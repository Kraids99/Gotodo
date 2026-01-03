package models

type Gotodo struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}