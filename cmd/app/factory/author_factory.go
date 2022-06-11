package factory

import (
	"learn-rest-api/cmd/app/form"
	"learn-rest-api/cmd/app/model"
)

func AuthorModelFactory(f form.AuthorForm) model.Author {
	return model.Author{
		Name: 			f.Name,
		IdentityNumber: f.IdentityNumber,
	}
}