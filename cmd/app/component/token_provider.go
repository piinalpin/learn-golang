package component

import (
	"learn-rest-api/cmd/app/repository"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
	Roles string
}

type Token struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    int64
}

type tokenProvider struct {
	userRepo repository.UserRepository
}

type TokenProvider interface {
	GenerateAccessToken(username string) Token
	IsTokenExpired(claims *JwtCustomClaims)
	ValidateAccessToken(token string) (claims *JwtCustomClaims, err string)
}

func TokenProviderInit(r repository.UserRepository) TokenProvider {
	godotenv.Load()
	return &tokenProvider{
		userRepo: r,
	}
}

var JWT_SIGNED_KEY string = os.Getenv("jwt.signed-key")

// GenerateAccessToken implements TokenProvider
func (t *tokenProvider) GenerateAccessToken(username string) Token {
	claims := &JwtCustomClaims{
		Roles: "ROLE_ADMIN",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &JwtCustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWT_SIGNED_KEY))
	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JWT_SIGNED_KEY))

	if err != nil {
		log.Error("Got error when generate access token. Error: ", err)
		return Token{}
	}

	return Token{
		AccessToken: token,
		RefreshToken: refreshToken,
		ExpiresAt: claims.StandardClaims.ExpiresAt,
	}
}

// IsTokenExpired implements TokenProvider
func (t *tokenProvider) IsTokenExpired(claims *JwtCustomClaims) {
	panic("unimplemented")
}

// ValidateAccessToken implements TokenProvider
func (t *tokenProvider) ValidateAccessToken(token string) (claims *JwtCustomClaims, err string) {
	panic("unimplemented")
}