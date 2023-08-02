package requests

type UserUpdateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	//	Username string `json:"username"`
	//	Profile string `json:"profile"`
}

type UserCreateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	//	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
