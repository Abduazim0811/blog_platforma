package postgres

import (
	"blog/internal/entity/posts"
	"blog/internal/entity/users"
	"blog/internal/infrastructura/repository"
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
)

type BlogsPostgres struct {
	db *sql.DB
}

func NewBlogsPostgres(db *sql.DB) repository.BlogsRepository {
	return &BlogsPostgres{db: db} 
}

func (b *BlogsPostgres) AddUsers(req users.CreateUserRequest) error {
	fmt.Println(req.Email, req.Password, req.UserName)
	sql, args, err := squirrel.
		Insert("users_blog").
		Columns("user_name, email, passwoord").
		Values(req.UserName, req.Email, req.Password).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("error generating SQL for AddUser:", err)
		return fmt.Errorf("error generating SQL for AddUser: %v", err)
	}

	_, err = b.db.Exec(sql, args...)
	if err != nil {
		log.Println("error exec add users:", err)
		return fmt.Errorf("error exec add users: %v", err)
	}

	return nil
}

func (b *BlogsPostgres) GetbyEmail(email string) (*users.Users, error) {
	var res users.Users
	sql, args, err := squirrel.
		Select("*").
		From("users_blog").
		Where(squirrel.Eq{"email": email}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("email not found")
		return nil, fmt.Errorf("email not found")
	}

	row := b.db.QueryRow(sql, args...)
	if err := row.Scan(&res.ID, &res.UserName, &res.Email, &res.Password); err != nil {
		log.Println("scan error")
		return nil, fmt.Errorf("scan error: %v", err)
	}

	return &res, nil
}

func (b *BlogsPostgres) GetUserByID(id int) (*users.Users, error) {
	sql, args, err := squirrel.
		Select("id, user_name, email, password").
		From("users_blog").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("error generating SQL for GetUserByID:", err)
		return nil, fmt.Errorf("error generating SQL for GetUserByID: %v", err)
	}

	var user users.Users
	err = b.db.QueryRow(sql, args...).Scan(&user.ID, &user.UserName, &user.Email, &user.Password)
	if err != nil {
		log.Println("error exec get user by ID:", err)
		return nil, fmt.Errorf("error exec get user by ID: %v", err)
	}

	return &user, nil
}

func (b *BlogsPostgres) PatchUpdateUser(user users.Users) error {
	query := squirrel.Update("users_blog").PlaceholderFormat(squirrel.Dollar)

	if user.UserName != "" {
		query = query.Set("user_name", user.UserName)
	}

	if user.Email != "" {
		query = query.Set("email", user.Email)
	}

	if user.Password != "" {
		query = query.Set("password", user.Password)
	}

	query = query.Where(squirrel.Eq{"id": user.ID})

	sql, args, err := query.ToSql()
	if err != nil {
		log.Println("error generating SQL for PatchUpdateUser:", err)
		return fmt.Errorf("error generating SQL for PatchUpdateUser: %v", err)
	}

	_, err = b.db.Exec(sql, args...)
	if err != nil {
		log.Println("error exec patch update user:", err)
		return fmt.Errorf("error exec patch update user: %v", err)
	}

	return nil
}

func (b *BlogsPostgres) DeleteUser(id int) error {
	var exists bool
	err := b.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users_blog WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		log.Println("error checking if user exists:", err)
		return fmt.Errorf("error checking if user exists: %v", err)
	}

	if !exists {
		return fmt.Errorf("user with id %d does not exist", id)
	}

	sql, args, err := squirrel.
		Delete("users_blog").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("error generating SQL for DeleteUser:", err)
		return fmt.Errorf("error generating SQL for DeleteUser: %v", err)
	}

	_, err = b.db.Exec(sql, args...)
	if err != nil {
		log.Println("error executing delete for user:", err)
		return fmt.Errorf("error executing delete for user: %v", err)
	}

	return nil
}

func (b *BlogsPostgres) AddPost(req posts.CreatePostRequest) (*posts.CreatePostResponse, error) {
	sql, args, err := squirrel.
		Insert("posts").
		Columns("user_id, title, content").
		Values(req.UserID, req.Title, req.Content).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("error generating SQL for AddPost:", err)
		return nil, fmt.Errorf("error generating SQL for AddPost: %v", err)
	}

	var postID int
	err = b.db.QueryRow(sql, args...).Scan(&postID)
	if err != nil {
		log.Println("error exec add post:", err)
		return nil, fmt.Errorf("error exec add post: %v", err)
	}

	return &posts.CreatePostResponse{ID: postID}, nil
}

func (b *BlogsPostgres) GetPostByID(id int) (*posts.Posts, error) {
	sql, args, err := squirrel.
		Select("id, user_id, title, content").
		From("posts").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("error generating SQL for GetPostByID:", err)
		return nil, fmt.Errorf("error generating SQL for GetPostByID: %v", err)
	}

	var post posts.Posts
	err = b.db.QueryRow(sql, args...).Scan(&post.ID, &post.UserID, &post.Title, &post.Content)
	if err != nil {
		log.Println("error exec get post by ID:", err)
		return nil, fmt.Errorf("error exec get post by ID: %v", err)
	}

	return &post, nil
}

func (b *BlogsPostgres) GetPostsByUserID(userID int) (*[]posts.Posts, error) {
	sqlPosts, argsPosts, err := squirrel.
		Select("id", "user_id", "title", "content").
		From("posts").
		Where(squirrel.Eq{"user_id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Println("error generating SQL for GetPostsByUserID:", err)
		return nil, fmt.Errorf("error generating SQL for GetPostsByUserID: %v", err)
	}

	rows, err := b.db.Query(sqlPosts, argsPosts...)
	if err != nil {
		log.Println("error fetching posts for user:", err)
		return nil, fmt.Errorf("error fetching posts for user: %v", err)
	}
	defer rows.Close()

	var userPosts []posts.Posts
	for rows.Next() {
		var post posts.Posts
		err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content)
		if err != nil {
			log.Println("error scanning post data:", err)
			return nil, fmt.Errorf("error scanning post data: %v", err)
		}
		userPosts = append(userPosts, post)
	}

	if err = rows.Err(); err != nil {
		log.Println("error after scanning posts:", err)
		return nil, fmt.Errorf("error after scanning posts: %v", err)
	}

	return &userPosts, nil
}

func (b *BlogsPostgres) PatchUpdatePost(post posts.Posts) error {
	query := squirrel.Update("posts").PlaceholderFormat(squirrel.Dollar)

	if post.Title != "" {
		query = query.Set("title", post.Title)
	}

	if post.Content != "" {
		query = query.Set("content", post.Content)
	}

	query = query.Where(squirrel.Eq{"id": post.ID})

	sql, args, err := query.ToSql()
	if err != nil {
		log.Println("error generating SQL for PatchUpdatePost:", err)
		return fmt.Errorf("error generating SQL for PatchUpdatePost: %v", err)
	}

	_, err = b.db.Exec(sql, args...)
	if err != nil {
		log.Println("error executing patch update post:", err)
		return fmt.Errorf("error executing patch update post: %v", err)
	}

	return nil
}

func (b *BlogsPostgres) DeletePost(id int) error {
	sql, args, err := squirrel.
		Delete("posts").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		log.Println("error generating SQL for DeletePost:", err)
		return fmt.Errorf("error generating SQL for DeletePost: %v", err)
	}

	_, err = b.db.Exec(sql, args...)
	if err != nil {
		log.Println("error executing delete post:", err)
		return fmt.Errorf("error executing delete post: %v", err)
	}

	return nil
}
