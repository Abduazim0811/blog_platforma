package handler

import (
	"blog/internal/entity/posts"
	"blog/internal/entity/users"
	jwt "blog/internal/pkg/token"
	"blog/internal/service"
	_ "blog/docs"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PostHandler struct {
	service service.PostsService
}

func NewPostsHandler(s service.PostsService) *PostHandler {
	return &PostHandler{service: s}
}

// @title Hotel Booking System
// @version 1.0
// @description This is a sample server for a Blogs reservation system.
// @securityDefinitions.apikey Bearer
// @in 				header
// @name Authorization
// @description Enter the token in the format `Bearer {token}`
// @host localhost:7777
// @BasePath /


// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param user body users.CreateUserRequest true "User request body"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /register [post]
func (p *PostHandler) Register(c *gin.Context) {
	var req users.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("error hashing password: %v", err)})
		return
	}
	req.Password = string(bytes)

	err = p.service.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "User Created"})
}

// Login godoc
// @Summary Login a user
// @Description Login a user and get a JWT token
// @Tags user
// @Accept json
// @Produce json
// @Param login body users.Login true "Login request body"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /login [post]
func (p *PostHandler) Login(c *gin.Context) {
	var req users.Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := p.service.GetUserByemail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect password"})
		return
	}

	token, err := jwt.GenerateJWTToken(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetByIdUser godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags user
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "User ID"
// @Success 200 {object} users.Users
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /users/{id} [get]
func (p *PostHandler) GetbyUserId(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := p.service.GetUserbyID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUsers godoc
// @Summary Update a user
// @Description Update a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body users.Users true "User request body"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users/{id} [put]
func (p *PostHandler) UpdateUser(c *gin.Context) {
	var req users.Users
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := p.service.Updateuser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated"})
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer
// @Router /users/{id} [delete]
func (p *PostHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err = p.service.Deleteuser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post
// @Tags post
// @Accept json
// @Produce json
// @Param post body posts.CreatePostRequest true "Post request body"
// @Success 200 {object} posts.CreatePostResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /posts [post]
func (p *PostHandler) CreatePost(c *gin.Context) {
	var req posts.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	resp, err := p.service.Createpost(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetPostByID godoc
// @Summary Get post by ID
// @Description Get post by ID
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} posts.Posts
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /posts/{id} [get]
func (p *PostHandler) GetPostByID(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	post, err := p.service.GetpostByID(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if post == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetPostsByUserID godoc
// @Summary Get posts by user ID
// @Description Get posts created by a specific user
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {array} posts.Posts
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /users/{id}/posts [get]
func (p *PostHandler) GetPostsByUserID(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	posts, err := p.service.GetPostsByuserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// UpdatePost godoc
// @Summary Update an existing post
// @Description Update an existing post
// @Tags post
// @Accept json
// @Produce json
// @Param post body posts.Posts true "Post request body"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /posts/{id} [put]
func (p *PostHandler) UpdatePost(c *gin.Context) {
	var req posts.Posts
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := p.service.Updatepost(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated"})
}

// DeletePost godoc
// @Summary Delete a post
// @Description Delete a post by ID
// @Tags post
// @Accept json
// @Produce json
// @Param id path string true "Post ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Security Bearer Auth
// @Router /posts/{id} [delete]
func (p *PostHandler) DeletePost(c *gin.Context) {
	id := c.Param("id")
	postID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post ID"})
		return
	}

	err = p.service.Deletepost(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}