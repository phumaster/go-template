package models

// User struct
type User struct {
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Model
}
