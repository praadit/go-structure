package entity

import (
	"time"

	"gorm.io/gorm"
)

type IBaseEntity[T any, K any] interface {
	ToResponse() (response K)
	ToUpdatable() (updatable map[string]interface{})
}

type BaseEntity struct {
	Id int64 `gorm:"primary_key; auto_increment;"`
}

type BaseTimeEntity struct {
	CreateDt time.Time `gorm:"autoCreateTime"`
	UpdateDt time.Time `gorm:"autoUpdateTime"`
	DeleteDt *gorm.DeletedAt
}
