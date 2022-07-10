package form

type RefreshTokenForm struct {
	RefreshToken string `json:"refresh_token" binding:"required,min=3"`
}