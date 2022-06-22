package component

import (
	respkey "learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/exception"
	"learn-rest-api/cmd/app/repository"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
	Roles string `json:"roles"`
}

type TokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type tokenProvider struct {
	userRepo repository.UserRepository
}

type TokenProvider interface {
	GenerateAccessToken(username string) TokenData
	ValidateAccessToken(accessToken string)
}

var JWT_SIGNED_KEY string
var JWT_ISSUER string

func TokenProviderInit(r repository.UserRepository) TokenProvider {
	godotenv.Load()
	JWT_SIGNED_KEY = os.Getenv("jwt.signed-key")
	JWT_ISSUER = os.Getenv("APPLICATION_NAME")
	return &tokenProvider{
		userRepo: r,
	}
}

// GenerateAccessToken implements TokenProvider
func (t *tokenProvider) GenerateAccessToken(username string) TokenData {
	claims := &JwtCustomClaims{
		Roles: "ROLE_ADMIN",
		StandardClaims: jwt.StandardClaims{
			Issuer:    JWT_ISSUER,
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    JWT_ISSUER,
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWT_SIGNED_KEY))
	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JWT_SIGNED_KEY))

	if err != nil {
		log.Error("Got error when generate access token. Error: ", err)
		return TokenData{}
	}

	return TokenData{
		AccessToken:  token,
		RefreshToken: refreshToken,
		ExpiresAt:    claims.StandardClaims.ExpiresAt,
	}
}

// ValidateAccessToken implements TokenProvider
func (t *tokenProvider) ValidateAccessToken(accessToken string) {
	authentication, err := jwt.ParseWithClaims(accessToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SIGNED_KEY), nil
	})

	if err != nil {
		log.Error("Got error when parse access token. ", err)
		exception.ThrowNewAppException(respkey.Unauthorized)
	}

	claims, ok := authentication.Claims.(*JwtCustomClaims)
	if !ok {
		exception.ThrowNewAppException_(respkey.Unauthorized.GetKey(), "Access token is invalid")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		exception.ThrowNewAppException_(respkey.Unauthorized.GetKey(), "Access token is expired")
	}
}
