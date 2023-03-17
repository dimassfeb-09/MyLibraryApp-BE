package request

type Category struct {
	ID   int    `json:"id"`
	Name string `binding:"required" json:"name"`
}
