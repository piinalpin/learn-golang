package component

import (
	respkey "learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/exception"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func Encode(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Error("Failed generate hash password.")
		exception.ThrowNewAppException(respkey.UnknownError)
	}

	return string(bytes)
}

func Matches(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		log.Error("Failed compare hash password.")
		return false
	}

	return true
}