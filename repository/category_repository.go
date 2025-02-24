package repository

import "amartha-go-struct/entity"

// interface = tempat declare method
type CategoryRepository interface {
	Create(category entity.CategoryEntity) (entity.CategoryEntity, error)
	FindById(id uint64) (entity.CategoryEntity, error)
	FindAll() ([]entity.CategoryEntity, error)
	Update(category entity.CategoryEntity) (entity.CategoryEntity, error)
}

// implement interface
type CategoryRepositoryImpl struct {
	categoryRepository CategoryRepository
}

func (repo *CategoryRepositoryImpl) Create(category entity.CategoryEntity) (entity.CategoryEntity, error) {
	return repo.categoryRepository.Create(category)
}

func (repo *CategoryRepositoryImpl) FindById(id uint64) (entity.CategoryEntity, error) {
	return repo.categoryRepository.FindById(id)
}

func (repo *CategoryRepositoryImpl) FindAll() ([]entity.CategoryEntity, error) {
	return repo.categoryRepository.FindAll()
}

func (repo *CategoryRepositoryImpl) Update(category entity.CategoryEntity) (entity.CategoryEntity, error) {
	return repo.categoryRepository.Update(category)
}

func NewCategoryRepositoryImpl(categoryRepository CategoryRepository) CategoryRepository {
	return &CategoryRepositoryImpl{}
}
