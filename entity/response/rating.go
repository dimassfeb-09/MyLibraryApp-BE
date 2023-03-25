package response

type Rating struct {
	ID     int     `json:"id"`
	Rating float32 `json:"rating"`
	BookID int     `json:"book_id"`
	UserID int     `json:"user_id"`
}

type RatingByBook struct {
	ID     int     `json:"id"`
	Rating float32 `json:"rating"`
	BookID int     `json:"book_id"`
}
