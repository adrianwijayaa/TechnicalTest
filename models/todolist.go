package models

type ToDoList struct {
	Id          int    `json:"id" gorm:"primaryKey:autoIncrement"`
	Title       string `json:"title" gorm:"type: varchar(100)"`
	Description string `json:"description" gorm:"type varchar(1000)"`
}
