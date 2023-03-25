package request

type Book struct {
	ID          int     `json:"id"`
	Title       string  `binding:"required" json:"title"`
	Description string  `binding:"required" json:"description"`
	Stok        int     `binding:"required" json:"stok"`
	Writer      string  `binding:"required" json:"writer"`
	ImgURL      string  `binding:"required" json:"img_url"`
	Rating      float32 `binding:"required" json:"rating"`
	CategoryID  int     `binding:"required" json:"category_id"`
	GenreID     int     `binding:"required" json:"genre_id"`
}
