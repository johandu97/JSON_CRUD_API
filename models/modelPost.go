package models

import "time"

// Post represents the Post table
type Post struct {
	ID        uint      `gorm:"primaryKey;auto-incrementing"`
	Title     string    `gorm:"type:varchar(255);not null" json:"Title" binding:"required"`
	Content   string    `gorm:"type:text;not null" json:"Content" binding:"required"`
	Author    string    `gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// MessageResponse represents a generic success or error message
type MessageResponse200 struct {
	Message string `json:"Message" example:"OK"`
}

type MessageResponse400 struct {
	Message string `json:"Message" example:"Bad Request"`
}

type MessageResponse500 struct {
	Message string `json:"Message" example:"Internal Server Error"`
}

// CreatePostRequest represents the JSON input for creating a post
type CreatePostRequest struct {
	Title   string `json:"Title" example:"My First Post"`        // Post Title
	Content string `json:"Content" example:"This is my content"` // Post Content
	Author  string `json:"Author" example:"John Doe"`            // Post Author
}
