package repository

import "amartha-go-struct/entity"

// interface = tempat declare method
type CategoryRepository interface {
	Create(category entity.Category) (entity.Category, error)
	FindById(id uint64) (entity.Category, error)
	FindAll() ([]entity.Category, error)
	Update(category entity.Category) (entity.Category, error)
}

// implement interface
type CategoryRepositoryImpl struct {
	categoryRepository CategoryRepository
}

func (repo *CategoryRepositoryImpl) Create(category entity.Category) (entity.Category, error) {
	return repo.categoryRepository.Create(category)
}

func (repo *CategoryRepositoryImpl) FindById(id uint64) (entity.Category, error) {
	return repo.categoryRepository.FindById(id)
}

func (repo *CategoryRepositoryImpl) FindAll() ([]entity.Category, error) {
	return repo.categoryRepository.FindAll()
}

func (repo *CategoryRepositoryImpl) Update(category entity.Category) (entity.Category, error) {
	return repo.categoryRepository.Update(category)
}

func NewCategoryRepositoryImpl(categoryRepository CategoryRepository) CategoryRepository {
	return &CategoryRepositoryImpl{}
}
