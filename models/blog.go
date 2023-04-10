package models

type Blog struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Image       string `json:"image"`
	UserId      string `json:"userId"`
	User        User   `json:"user";gorm:"foreignkey:UserId"`
}
