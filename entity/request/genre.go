package request

type Genre struct {
	ID         int    `json:"id"`
	Name       string `binding:"required" json:"name"`
	CategoryID int    `binding:"required" json:"category_id"`
}
