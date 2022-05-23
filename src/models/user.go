package models

type User struct {
	Name  string `json:"name"`
	Email string `gorm:"primary_key" json:"email"`
}
