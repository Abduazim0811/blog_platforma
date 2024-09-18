package api

import (
	"blog/internal/https/handler"
	"blog/internal/https/middleware"
	"blog/internal/infrastructura/postgres"
	"blog/internal/service"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(db *sql.DB) {
	repo := postgres.NewBlogsPostgres(db)
	service := service.NewPostsService(repo)
	handlers := handler.NewPostsHandler(*service)

	r := gin.Default()
	// users
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/users/:id", middleware.Protected(), handlers.GetbyUserId)
	r.PUT("/users/:id", middleware.Protected(), handlers.UpdateUser)
	r.DELETE("/users/:id", middleware.Protected(), handlers.DeleteUser)
	// posts
	r.POST("/posts", middleware.Protected(), handlers.CreatePost)
	r.GET("/posts/:id", middleware.Protected(), handlers.GetPostByID)
	r.GET("/users/:id/posts", middleware.Protected(), handlers.GetPostsByUserID)
	r.PUT("/posts/:id", middleware.Protected(), handlers.UpdatePost)
	r.DELETE("/posts/:id", middleware.Protected(), handlers.DeletePost)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.RunTLS(":7777", "./internal/tls/items.pem", "./internal/tls/items-key.pem")
	if err != nil {
		log.Fatal("Failed to run HTTPS server:", err)
	}
}
