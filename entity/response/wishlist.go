package response

type Wishlist struct {
	ID     int `json:"id"`
	BookID int `binding:"required" json:"book_id"`
	UserID int `binding:"required" json:"user_id"`
}
