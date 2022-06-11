package form

type AuthorForm struct {
	Name 			string `json:"name" binding:"required,min=3,max=50"`
	IdentityNumber	string `json:"identity_number" binding:"required,len=16"`
}
