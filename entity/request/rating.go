package request

type Rating struct {
	ID     int     `json:"id"`
	Rating float32 `binding:"required" json:"rating"`
	BookID int     `binding:"required" json:"book_id"`
	UserID int     `binding:"required" json:"user_id"`
}
