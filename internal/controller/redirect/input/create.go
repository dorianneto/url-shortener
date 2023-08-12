package input

type CreateRedirect struct {
	Url string `json:"url" binding:"required"`
}
