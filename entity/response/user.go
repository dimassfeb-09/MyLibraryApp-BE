package response

type UserResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	NPM      string `json:"npm"`
	Email    string `json:"email"`
	IsGoogle bool   `json:"is_google"`
}
