package main

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
	"fmt"
	"log"
	"time"
)

func main() {

	initializers.ConnectToDB()

	// Retrieve all posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		log.Fatalf("Failed to retrieve posts: %v", result.Error)
	}

	// Display posts
	fmt.Println("Posts:")
	for _, post := range posts {
		fmt.Printf("ID: %d\nTitle: %s\nContent: %s\nAuthor: %s\nCreated At: %s\nUpdated At: %s\n\n",
			post.ID, post.Title, post.Content, post.Author, post.CreatedAt.Format(time.RFC3339), post.UpdatedAt.Format(time.RFC3339))
	}
}
