package posts


type CreatePostRequest struct {
    UserID  int `json:"user_id"`  
    Title   string `json:"title"` 
    Content string `json:"content"` 
}

type CreatePostResponse struct{
	ID		int	`json:"id"`
}

type Posts struct{
	ID		int	`json:"id"`
	UserID  int `json:"user_id"`  
    Title   string `json:"title"` 
    Content string `json:"content"` 
}