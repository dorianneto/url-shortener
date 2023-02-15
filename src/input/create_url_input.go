package input

type CreateUrlInput struct {
	Url string `json:"url" binding:"required"`
}
