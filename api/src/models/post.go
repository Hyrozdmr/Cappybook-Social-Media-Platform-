package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	Likes     int       `json:"likes"`
}

func (post *Post) Save() (*Post, error) {
	err := Database.Create(post).Error
	if err != nil {
		return &Post{}, err
	}

	return post, nil
}

func (post *Post) SaveLike() (*Post, error) {
	post.Likes++
	err := Database.Save(post).Error
	if err != nil {
		return nil, err
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

// func FetchSinglePost() (*[]Post, error) {
// 	var posts []Post
// 	err := Database.Find(&posts).Error

// 	fmt.Println(posts)

// 	if err != nil {
// 		return &[]Post{}, err
// 	}

// 	return &posts, nil
// }

func FetchSpecificPost(postID uint64) (*Post, error) {
	var post Post
	err := Database.First(&post, postID).Error

	if err != nil {
		return nil, err
	}

	return &post, nil
}
