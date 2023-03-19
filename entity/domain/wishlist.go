package domain

type Wishlist struct {
	ID     int `json:"id"`
	BookID int `json:"book_id"`
	UserID int `json:"user_id"`
}
