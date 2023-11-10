package types

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type Posts struct {
	gorm.Model
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"user_id"`
	User   Users  `gorm:"foreignKey:UserID;references:ID" json:"user"`
}
