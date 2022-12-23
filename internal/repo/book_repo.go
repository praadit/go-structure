package repos

import (
	"go-best-practice/internal/entity"
	"go-best-practice/internal/request"

	"gorm.io/gorm"
)

type IBookRepo interface {
	GetAll(db *gorm.DB) (list []entity.Book, err error)
	GetPaged(db *gorm.DB, pageReq request.PaginationRequest) (list []request.BookData, total int64, err error)
	GetById(db *gorm.DB, id int64) (data entity.Book, err error)
	Create(db *gorm.DB, newEntity entity.Book) (err error)
	Update(db *gorm.DB, updatedEntity entity.Book) (err error)
	Delete(db *gorm.DB, id int64) (err error)
}

type BookRepo struct {
	baseRepo BaseRepository[entity.Book]
}

func CreateBookRepo() IBookRepo {
	en := entity.Book{}
	baseRepo := BaseRepository[entity.Book]{
		model:      en,
		tableName:  "book",
		orderable:  en.ToOrderable(),
		searchable: en.ToSearchable(),
	}
	return &BookRepo{
		baseRepo: baseRepo,
	}
}

func (repo *BookRepo) GetAll(db *gorm.DB) (list []entity.Book, err error) {
	tables := repo.baseRepo.GetAll(db)
	err = tables.Find(&list).Error
	return
}

func (repo *BookRepo) GetPaged(db *gorm.DB, pageReq request.PaginationRequest) (list []request.BookData, total int64, err error) {
	data, total, err := repo.baseRepo.GetPagedData(db, pageReq)
	list = []request.BookData{}
	for _, v := range data {
		list = append(list, *v.ToResponse())
	}
	return
}

func (repo *BookRepo) GetById(db *gorm.DB, id int64) (data entity.Book, err error) {
	err = repo.baseRepo.GetById(db, id, &data)
	return
}

func (repo *BookRepo) Create(db *gorm.DB, newEntity entity.Book) (err error) {
	err = repo.baseRepo.Create(db, newEntity)
	return
}

func (repo *BookRepo) Update(db *gorm.DB, updatedEntity entity.Book) (err error) {
	updatable := updatedEntity.ToUpdatable()
	err = repo.baseRepo.Update(db, updatedEntity.Id, updatable)

	return
}

func (repo *BookRepo) Delete(db *gorm.DB, id int64) (err error) {
	err = repo.baseRepo.Delete(db, id)
	return
}
