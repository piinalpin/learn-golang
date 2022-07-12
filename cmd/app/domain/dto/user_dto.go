package dto

import "learn-rest-api/cmd/app/domain/base"

type UserDto struct {
	base.BaseDto
	ID 				uint 		`json:"id"`
	Username		string		`json:"username"`
	Status			string		`json:"status"`
}