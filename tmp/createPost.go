package main

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
	"fmt"
	"log"
)

func main() {

	initializers.ConnectToDB()

	// Create a new post
	newPost := models.Post{
		Title:   "My First Post",
		Content: "This is the content of my first post.",
		Author:  "John Doe",
	}

	result := initializers.DB.Create(&newPost)
	if result.Error != nil {
		log.Fatalf("Failed to create post: %v", result.Error)
	}

	fmt.Printf("Post created successfully with ID: %d\n", newPost.ID)
}
