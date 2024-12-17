package main

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

func main() {

	initializers.ConnectToDB()

	// Retrieve a single post by ID
	var post models.Post
	postID := 3 // Change this to the ID of the post you want to view
	result := initializers.DB.First(&post, postID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			log.Printf("Post with ID %d not found", postID)
		} else {
			log.Fatalf("Failed to retrieve post: %v", result.Error)
		}
		return
	}

	// Display the post
	fmt.Printf("Post Details:\n")
	fmt.Printf("ID: %d\nTitle: %s\nContent: %s\nAuthor: %s\nCreated At: %s\nUpdated At: %s\n",
		post.ID, post.Title, post.Content, post.Author, post.CreatedAt.Format(time.RFC3339), post.UpdatedAt.Format(time.RFC3339))
}
