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
	rawMap, err := utilities.ToMap(en)
	if err != nil {
		panic(err)
	}
	updatable = utilities.KeyToSnakeCase(rawMap)
	delete(updatable, "id")
	delete(updatable, "delete_dt")

	return
}
func (en *User) ToOrderable() (orderable []string) {
	rawMap, err := utilities.ToMap(en)
	if err != nil {
		panic(err)
	}
	fields := utilities.KeyToSnakeCase(rawMap)
	delete(fields, "password")
	delete(fields, "delete_dt")

	orderable = utilities.GetMapKeys(fields)

	return
}

func (en *User) ToSearchable() (searchable []string) {
	rawMap, err := utilities.ToMap(en)
	if err != nil {
		panic(err)
	}
	fields := utilities.KeyToSnakeCase(rawMap)
	delete(fields, "password")
	delete(fields, "created_at")
	delete(fields, "updated_at")
	delete(fields, "delete_dt")

	searchable = utilities.GetMapKeys(fields)

	return
}
