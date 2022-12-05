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
	updatable, err := utilities.ToMap(en)
	if err != nil {
		panic(err)
	}

	return
}
