package database

import (
	"fmt"
	"todo-app/models"
	"todo-app/pkg/mysql"
)

func AutoMigration() {
	err := mysql.DB.AutoMigrate(&models.ActivityGroup{}, &models.Todo{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("migration success")
}
