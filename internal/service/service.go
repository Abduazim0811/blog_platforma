package service

import (
	"blog/internal/entity/posts"
	"blog/internal/entity/users"
	"blog/internal/infrastructura/repository"
)

type PostsService struct {
	repo repository.BlogsRepository
}

func NewPostsService(repo repository.BlogsRepository) *PostsService {
	return &PostsService{repo: repo}
}

func (p *PostsService) CreateUser(req users.CreateUserRequest) error {
	return p.repo.AddUsers(req)
}

func (p *PostsService) GetUserbyID(id int) (*users.Users, error) {
	return p.repo.GetUserByID(id)
}

func (p *PostsService) GetUserByemail(email string) (*users.Users, error) {
	return p.repo.GetbyEmail(email)
}

func (p *PostsService) Updateuser(user users.Users) error {
	return p.repo.PatchUpdateUser(user)
}

func (p *PostsService) Deleteuser(id int) error {
	return p.repo.DeleteUser(id)
}

func (p *PostsService) Createpost(req posts.CreatePostRequest) (*posts.CreatePostResponse, error) {
	return p.repo.AddPost(req)
}

func (p *PostsService) GetpostByID(id int) (*posts.Posts, error) {
	return p.repo.GetPostByID(id)
}

func (p *PostsService) GetPostsByuserID(userID int) (*[]posts.Posts, error) {
	return p.repo.GetPostsByUserID(userID)
}

func (p *PostsService) Updatepost(post posts.Posts) error {
	return p.repo.PatchUpdatePost(post)
}

func (p *PostsService) Deletepost(id int) error {
	return p.repo.DeletePost(id)
}
