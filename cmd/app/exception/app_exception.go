package exception

import (
	"errors"
	"fmt"
	respkey "learn-rest-api/cmd/app/constant"
)

func ThrowNewAppException_(key string, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func ThrowNewAppException(responseKey respkey.ResponseKey) {
	ThrowNewAppException_(responseKey.GetKey(), responseKey.GetMessage())
}