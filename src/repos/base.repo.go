package repos

import (
	"errors"
	"fmt"
	"go-best-practice/src/models/dto"
	"go-best-practice/src/utilities"
	"strings"

	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	model      T
	tableName  string
	orderable  []string
	searchable []string
}

func (repo *BaseRepository[T]) GetPagedData(db *gorm.DB, pagReq dto.PaginationRequest) (entity []T, total int64, err error) {
	orderBy := strings.ToLower(pagReq.OrderBy)
	if !utilities.Contains(repo.orderable, orderBy) {
		err = gorm.ErrInvalidField
		return
	}

	validOrder := []string{"asc", "desc"}
	if !utilities.Contains(validOrder, strings.ToLower(pagReq.Order)) {
		err = gorm.ErrInvalidField
		return
	}

	search := "delete_dt is null "

	if pagReq.Search != "" {
		search += utilities.CreateSearchQuery(repo.searchable, pagReq.Search)
	}

	offset := pagReq.Perpage * (pagReq.Page + 1)

	queryTotal := fmt.Sprintf(`select count(id) as total from %s where %s`, repo.tableName, search)
	err = db.Raw(queryTotal, search).Find(&total).Error
	if err != nil {
		return
	}

	query := fmt.Sprintf(`select * from %s 
		where %s 
		order by %s %s
		offset %d
		limit %d
	`, repo.tableName, search, orderBy, pagReq.Order, offset, pagReq.Perpage)

	err = db.Raw(query).Find(&entity).Error
	return
}

func (repo *BaseRepository[T]) GetAll(db *gorm.DB) *gorm.DB {
	return db.Model(&repo.model)
}

func (repo *BaseRepository[T]) GetById(db *gorm.DB, id int64, data any) (err error) {
	return db.Model(&repo.model).Where("id", id).First(data).Error
}

func (repo *BaseRepository[T]) Create(db *gorm.DB, newData any) error {
	return db.Create(newData).Error
}

func (repo *BaseRepository[T]) Update(db *gorm.DB, id int64, changes map[string]any) error {
	return db.Model(&repo.model).Where("id", id).Updates(changes).Error
}

func (repo *BaseRepository[T]) Delete(db *gorm.DB, id int64) error {
	data := repo.model
	err := repo.GetById(db, id, &data)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return db.Delete(data).Error
}
