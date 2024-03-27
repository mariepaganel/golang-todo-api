package models

type Todo struct {
	ID          int    `json:"id" gorm:"primaryKey; unique"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
