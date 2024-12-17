package main

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
	"log"
)

func main() {

	initializers.ConnectToDB()

	// Migrate the schema to create the posts table (auto-create the table based on the Post struct)
	err := initializers.DB.AutoMigrate(&models.Post{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Table 'posts' created successfully (if it doesn't already exist).")
}
