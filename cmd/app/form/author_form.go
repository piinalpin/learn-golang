package form

type AuthorForm struct {
	Name string `json:"name" binding:"required,min=3,max=50"`
}
