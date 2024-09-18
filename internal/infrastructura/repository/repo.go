package repository

import (
	"blog/internal/entity/posts"
	"blog/internal/entity/users"
)

type BlogsRepository interface {
	AddUsers(req users.CreateUserRequest) error
	GetbyEmail(email string)(*users.Users, error)
	GetUserByID(id int) (*users.Users, error)
	PatchUpdateUser(user users.Users) error
	DeleteUser(id int) error
	AddPost(req posts.CreatePostRequest) (*posts.CreatePostResponse, error)
	GetPostByID(id int) (*posts.Posts, error)
	GetPostsByUserID(userID int) (*[]posts.Posts, error)
	PatchUpdatePost(post posts.Posts) error
	DeletePost(id int) error
}
