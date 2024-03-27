package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

func (post *Post) Save() (*Post, error) {
	err := Database.Create(post).Error
	if err != nil {
		return &Post{}, err
	}

	return post, nil
}

func FetchAllPosts() (*[]Post, error) {
	var posts []Post
	err := Database.Find(&posts).Error

	fmt.Println(posts)

	if err != nil {
		return &[]Post{}, err
	}

	return &posts, nil
}
