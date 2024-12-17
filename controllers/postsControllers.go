package controllers

import (
	"JSON_CRUD_API/initializers"
	"JSON_CRUD_API/models"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPost godoc
// @Summary Create a table
// @Description Create a table with name "posts"
// @Tags Process Table
// @Produce json
// @Success 200 {object} models.MessageResponse200 "OK"
// @Failure 400 {object} models.MessageResponse400 "Bad Request"
// @Router /createtable [post]
func CreateTable(c *gin.Context) {

	// Migrate the schema to create the posts table (auto-create the table based on the Post struct)
	err := initializers.DB.AutoMigrate(&models.Post{})

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Create Failure",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Create Successfully",
	})

}

// GetPost godoc
// @Summary Create a post
// @Description Create a post with content, title, author
// @Tags Process Post
// @Produce json
// @Param post body models.CreatePostRequest true "Create Post"
// @Success 200 {object} models.MessageResponse200 "OK"
// @Failure 400 {object} models.MessageResponse400 "Bad Request"
// @Failure 500 {object} models.MessageResponse500 "Internal Server Error"
// @Router /createpost [post]
func CreatePost(c *gin.Context) {

	var newPost models.Post

	// Parse and bind JSON to struct
	if err := c.ShouldBindJSON(&newPost); err != nil {
		// Handle error if JSON is invalid or fields are missing
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	result := initializers.DB.Create(&newPost)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		return
	}

	fmt.Printf("Post created successfully with ID: %d\n", newPost.ID)

	// Respond with the parsed data
	c.JSON(http.StatusOK, gin.H{
		"Message": "OK",
		"ID":      newPost.ID,
	})
}

// GetPost godoc
// @Summary View a post
// @Description View a post with ID
// @Tags Process Post
// @Produce json
// @Param ID path int true "Post ID"  // Define the path parameter as an integer
// @Success 200 {object} models.MessageResponse200 "OK"
// @Failure 400 {object} models.MessageResponse400 "Bad Request"
// @Router /viewpost/{ID} [get]
func ViewPost(c *gin.Context) {

	var post models.Post

	// Extract the ID from the path parameter
	idParam := c.Param("ID")

	postID := idParam // Change this to the ID of the post you want to view
	result := initializers.DB.First(&post, postID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		if result.Error == gorm.ErrRecordNotFound {
			fmt.Println("Post with ID %d not found", postID)
		} else {
			fmt.Println("Failed to retrieve post: %v", result.Error)
		}
		return
	}

	// Format the response
	response := gin.H{
		"id":         post.ID,
		"title":      post.Title,
		"content":    post.Content,
		"author":     post.Author,
		"created_at": post.CreatedAt.Format(time.RFC3339), // Format time to RFC3339
		"updated_at": post.UpdatedAt.Format(time.RFC3339), // Format time to RFC3339
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"Message": "OK",
		"Data":    response,
	})

}

// GetPost godoc
// @Summary Update a post
// @Description Update a post with ID
// @Tags Process Post
// @Produce json
// @Param ID path int true "Post ID"  // Define the path parameter as an integer
// @Param post body models.CreatePostRequest true "Update Post"
// @Success 200 {object} models.MessageResponse200 "OK"
// @Failure 400 {object} models.MessageResponse400 "Bad Request"
// @Failure 500 {object} models.MessageResponse500 "Internal Server Error"
// @Router /updatepost/{ID} [put]
func UpdatePost(c *gin.Context) {

	var newPost models.Post

	// Parse and bind JSON to struct
	if err := c.ShouldBindJSON(&newPost); err != nil {
		// Handle error if JSON is invalid or fields are missing
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	var post models.Post

	// Extract the ID from the path parameter
	postID := c.Param("ID")

	// Find the post by ID
	if err := initializers.DB.First(&post, postID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		if err == gorm.ErrRecordNotFound {
			log.Printf("Post with ID %d not found", postID)
		} else {
			log.Fatalf("Failed to retrieve post: %v", err)
		}
		return
	}

	// Modify fields
	post.Title = newPost.Title
	post.Content = newPost.Content
	post.Author = newPost.Author

	fmt.Println(post)

	// Save the changes
	if err := initializers.DB.Save(&post).Error; err != nil {
		log.Fatalf("Failed to update post: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "OK",
		})
		log.Printf("Post with ID %d updated successfully", postID)
	}

}

// GetPost godoc
// @Summary Delete a post
// @Description Delete a post with ID
// @Tags Process Post
// @Produce json
// @Param ID path int true "Post ID"  // Define the path parameter as an integer
// @Success 200 {object} models.MessageResponse200 "OK"
// @Failure 400 {object} models.MessageResponse400 "Bad Request"
// @Router /deletepost/{ID} [delete]
func DeletePost(c *gin.Context) {

	// Extract the ID from the path parameter
	postID := c.Param("ID")

	// Delete post with ID 1
	result := initializers.DB.Delete(&models.Post{}, postID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		log.Fatalf("Failed to delete post: %v", result.Error)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Message": "OK",
		})
		log.Printf("Post with ID %d deleted successfully", postID)
	}
}

// GetPost godoc
// @Summary View many posts
// @Description View many posts
// @Tags Process Post
// @Produce json
// @Success 200 {object} models.MessageResponse200 "OK"
// @Failure 400 {object} models.MessageResponse400 "Bad Request"
// @Router /viewposts [get]
func ViewPosts(c *gin.Context) {

	// Retrieve all posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Bad Request",
		})
		log.Fatalf("Failed to retrieve posts: %v", result.Error)
		return
	}

	// Format the response: Map each post to a simplified structure
	var formattedPosts []map[string]interface{}
	for _, post := range posts {
		formattedPosts = append(formattedPosts, map[string]interface{}{
			"id":         post.ID,
			"title":      post.Title,
			"content":    post.Content,
			"author":     post.Author,
			"created_at": post.CreatedAt.Format(time.RFC3339),
			"updated_at": post.UpdatedAt.Format(time.RFC3339),
		})
	}

	// Return the response with all posts
	c.JSON(http.StatusOK, gin.H{
		"Message": "OK",
		"Data":    formattedPosts,
	})
}
