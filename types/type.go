package types

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Posts struct {
	gorm.Model
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"user_id"`
	User   Users  `gorm:"foreignKey:UserID;references:ID" json:"user"`
}
