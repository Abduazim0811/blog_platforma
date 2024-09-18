package users

type Users struct{
	ID 		 int	`json:"id"`
	UserName string `json:"user_name"`  
    Email    string `json:"email"` 
    Password string `json:"passwoord"`
}
type CreateUserRequest struct {
    UserName string `json:"user_name"`  
    Email    string `json:"email"` 
    Password string `json:"password"` 
}

type Login struct{
	Email    string `json:"email"` 
    Password string `json:"password"` 
}

type GetUserRequest struct {
    ID string `json:"id"`
}