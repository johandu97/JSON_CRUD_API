package main

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
	"log"

	"gorm.io/gorm"
)

func main() {

	initializers.ConnectToDB()

	// Update post with ID 1
	postID := 1
	var post models.Post

	// Find the post by ID
	if err := initializers.DB.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Post with ID %d not found", postID)
		} else {
			log.Fatalf("Failed to retrieve post: %v", err)
		}
		return
	}

	// Find the post by ID
	if err := initializers.DB.First(&post, postID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("Post with ID %d not found", postID)
		} else {
			log.Fatalf("Failed to retrieve post: %v", err)
		}
		return
	}

	// Modify fields
	post.Title = "Updated Post Title"
	post.Content = "Updated content of the post."

	// Save the changes
	if err := initializers.DB.Save(&post).Error; err != nil {
		log.Fatalf("Failed to update post: %v", err)
	} else {
		log.Printf("Post with ID %d updated successfully", postID)
	}
}
