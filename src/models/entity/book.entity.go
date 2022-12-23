package entity

import (
	"go-best-practice/src/models/dto"
	"go-best-practice/src/utilities"
)

type Book struct {
	BaseEntity

	AuthorId int64
	BookName string

	BaseTimeEntity

	Author User `gorm:"foreign_key:author_id"`
}

func (en *Book) ToResponse() (response *dto.BookData) {
	response = &dto.BookData{
		BookName: en.BookName,
	}

	return
}

func (en *Book) ToUpdatable() (updatable map[string]interface{}) {
	rawMap, err := utilities.ToMap(en)
	if err != nil {
		panic(err)
	}
	updatable = utilities.KeyToSnakeCase(rawMap)
	delete(updatable, "id")
	delete(updatable, "delete_dt")

	return
}

func (en *Book) ToOrderable() (orderable []string) {
	rawMap, err := utilities.ToMap(en)
	if err != nil {
		panic(err)
	}
	fields := utilities.KeyToSnakeCase(rawMap)
	delete(fields, "author_id")
	delete(fields, "delete_dt")

	orderable = utilities.GetMapKeys(fields)

	return
}

func (en *Book) ToSearchable() (searchable []string) {
	rawMap, err := utilities.ToMap(en)
	if err != nil {
		panic(err)
	}
	fields := utilities.KeyToSnakeCase(rawMap)
	delete(fields, "created_at")
	delete(fields, "updated_at")
	delete(fields, "delete_dt")

	searchable = utilities.GetMapKeys(fields)

	return
}
