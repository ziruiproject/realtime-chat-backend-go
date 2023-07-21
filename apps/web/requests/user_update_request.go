package requests

type UserUpdateRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	//	Username string `json:"username"`
	Profile string `json:"profile"`
}
