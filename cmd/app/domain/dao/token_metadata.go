package dao

type TokenMetadata struct {
	User		User 		`json:"user"`
}

type RefreshTokenMetadata struct {
	Username 			string 		`json:"username"`
	AccessTokenUuid 	string 		`json:"access_token_uuid"`
}