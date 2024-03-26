package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model // The gorm model provides common fields like ID, CreatedAt, UpdatedAt and DeletedAt
	// **gorm.Model LIKELY TO BE USEFUL WHEN WE NEED TO ADD THE RELEVANT DATES**
	Message string `json:"message"`
}

// This function creates a new record in the database
func (post *Post) Save() (*Post, error) {
	err := Database.Create(post).Error // Database relates to the database connection variable created in the db.go file. Create and Error are part of the functions imported from gorm
	if err != nil {
		return &Post{}, err
	}

	return post, nil
}

// This function retrieves all posts from the database and returns them as a slice
func FetchAllPosts() (*[]Post, error) {
	var posts []Post
	err := Database.Find(&posts).Error // Database.Find(&posts) fetches all records from the database and adds them to the 'posts' slice

	fmt.Println(posts) // Prints out the slice (likely used for debugging)

	if err != nil {
		return &[]Post{}, err
	}

	return &posts, nil
}
