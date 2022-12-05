package entity

import (
	"go-best-practice/src/models/dto"
	"go-best-practice/src/utilities"
)

type User struct {
	BaseEntity

	Username string
	Password string

	BaseTimeEntity
}

func (en *User) ToResponse() (response *dto.UserData) {
	response = &dto.UserData{
		Username: en.Username,
		Password: en.Password,
	}

	return
}

func (en *User) ToUpdatable() (updatable map[string]interface{}) {
	updatable, err := utilities.ToMap(en)
	if err != nil {
		panic(err)
	}

	return
}
