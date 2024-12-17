package main

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
	"log"
)

func main() {

	initializers.ConnectToDB()

	// Delete post with ID 1
	postID := 2
	result := initializers.DB.Delete(&models.Post{}, postID)
	if result.Error != nil {
		log.Fatalf("Failed to delete post: %v", result.Error)
	} else {
		log.Printf("Post with ID %d deleted successfully", postID)
	}
}
