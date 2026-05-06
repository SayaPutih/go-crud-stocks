package requests

type UpdateUserRequest struct {
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	IsVerified  bool   `json:"is_verified"`
}
