package user

type User struct{
	Id          int    `json:"id"`
	Name       string `json:"name"`
	Email       string    `json:"email"`
	Password       []byte `json:"password"`
	Role       string    `json:"role"`
}

type UserService interface{
    GetUserTest() map[string]string
    GetUsers() ([]User,error) 
}

