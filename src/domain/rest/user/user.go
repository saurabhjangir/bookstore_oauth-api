package user

type UserLoginRequest struct{
	Password string `json:"password"`
	Email string `json:"email"`
}

type UserLoginResponse struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}