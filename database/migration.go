package database

import (
	"TechnicalTest/models"
	"TechnicalTest/pkg/postgresql"
	"fmt"
)

func RunMigration() {
	err := postgresql.DB.AutoMigrate(&models.ToDoList{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
