package input

type FindRedirect struct {
	Code string `uri:"code" binding:"required"`
}
