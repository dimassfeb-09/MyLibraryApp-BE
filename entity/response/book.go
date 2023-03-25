package response

type Book struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Stok        int     `json:"stok"`
	Writer      string  `json:"writer"`
	ImgURL      string  `json:"img_url"`
	Rating      float32 `json:"rating"`
	CategoryID  int     `json:"category_id"`
	GenreID     int     `json:"genre_id"`
}
