package component

import (
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/domain/dao"
	"learn-rest-api/cmd/app/exception"
	"learn-rest-api/cmd/app/repository"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
	Roles string `json:"roles"`
	Uuid  string `json:"uuid"`
}

type TokenData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type tokenProvider struct {
	userRepo       repository.UserRepository
	sessionStorage SessionStorage
}

type TokenProvider interface {
	GenerateAccessToken(user dao.User) TokenData
	ValidateAccessToken(accessToken string) *JwtCustomClaims
	RefreshToken(refreshToken string) TokenData
}

var JWT_SIGNED_KEY string
var JWT_ISSUER string
var ACCESS_TOKEN_LIFETIME int
var REFRESH_TOKEN_LIFETIME int

func TokenProviderInit(r repository.UserRepository, ss SessionStorage) TokenProvider {
	godotenv.Load()
	JWT_SIGNED_KEY = os.Getenv("jwt.signed-key")
	JWT_ISSUER = os.Getenv("application.name")
	ACCESS_TOKEN_LIFETIME, _ = strconv.Atoi(os.Getenv("jwt.access-token.lifetime"))
	REFRESH_TOKEN_LIFETIME, _ = strconv.Atoi(os.Getenv("jwt.refresh-token.lifetime"))
	return &tokenProvider{
		userRepo:       r,
		sessionStorage: ss,
	}
}

// GenerateAccessToken implements TokenProvider
func (t *tokenProvider) GenerateAccessToken(user dao.User) TokenData {
	var rolesArray []string

	for i := 0; i < len(user.Roles); i++ {
		userRole := user.Roles[i]
		rolesArray = append(rolesArray, userRole.Role)
	}

	roles := strings.Join(rolesArray, ",")
	claims := &JwtCustomClaims{
		Roles: roles,
		Uuid:  uuid.New().String(),
		StandardClaims: jwt.StandardClaims{
			Issuer:    JWT_ISSUER,
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(ACCESS_TOKEN_LIFETIME * 60)).Unix(),
		},
	}

	refreshClaims := &JwtCustomClaims{
		Uuid:  uuid.New().String(),
		StandardClaims: jwt.StandardClaims{
			Issuer:    JWT_ISSUER,
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(REFRESH_TOKEN_LIFETIME * 60)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWT_SIGNED_KEY))
	refreshToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JWT_SIGNED_KEY))

	if err != nil {
		log.Error("Got error when generate access token. Error: ", err)
		return TokenData{}
	}

	tokenMetadata := dao.TokenMetadata{
		User: user,
	}

	refreshTokenMetadata := dao.RefreshTokenMetadata{
		Username: user.Username,
		AccessTokenUuid: claims.Uuid,
	}

	log.Info("Store token metadata to cache")
	t.sessionStorage.PutCache(constant.UserSession.GetCacheName(), claims.Uuid, tokenMetadata, ACCESS_TOKEN_LIFETIME)
	t.sessionStorage.PutCache(constant.RefreshTokenSession.GetCacheName(), refreshClaims.Uuid, refreshTokenMetadata, REFRESH_TOKEN_LIFETIME)

	return TokenData{
		AccessToken:  token,
		RefreshToken: refreshToken,
		ExpiresAt:    claims.StandardClaims.ExpiresAt,
	}
}

// ValidateAccessToken implements TokenProvider
func (t *tokenProvider) ValidateAccessToken(accessToken string) *JwtCustomClaims {
	authentication, err := jwt.ParseWithClaims(accessToken, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_SIGNED_KEY), nil
	})

	if err != nil {
		log.Error("Got error when parse access token. ", err)
		exception.ThrowNewAppException(constant.Unauthorized)
	}

	claims, ok := authentication.Claims.(*JwtCustomClaims)
	if !ok {
		exception.ThrowNewAppException_(constant.Unauthorized.GetKey(), "Access token is invalid")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		exception.ThrowNewAppException_(constant.Unauthorized.GetKey(), "Access token is expired")
	}

	return claims
}

// RefreshToken implements TokenProvider
func (t *tokenProvider) RefreshToken(refreshToken string) TokenData {
	log.Info("Validating refresh token")
	claims := t.ValidateAccessToken(refreshToken)

	log.Info("Getting token metadata from cache")

	var refreshTokenMetadata dao.RefreshTokenMetadata
	err := t.sessionStorage.GetCache(constant.RefreshTokenSession.GetCacheName(), claims.Uuid, &refreshTokenMetadata)
	if err != nil {
		exception.ThrowNewAppException_(constant.Unauthorized.GetKey(), constant.Unauthorized.GetMessage())
	}

	log.Info("Delete existing token metadata from cache")
	t.sessionStorage.RemoveCache(constant.UserSession.GetCacheName(), refreshTokenMetadata.AccessTokenUuid)
	t.sessionStorage.RemoveCache(constant.RefreshTokenSession.GetCacheName(), claims.Uuid)

	log.Info("Getting user from database")
	user, err := t.userRepo.FindUserByUsername(refreshTokenMetadata.Username)

	if err != nil {
		log.Error("Got error when get user from database. Error: ", err)
		exception.ThrowNewAppException_(constant.Unauthorized.GetKey(), constant.Unauthorized.GetMessage())
	}

	return t.GenerateAccessToken(user)
}