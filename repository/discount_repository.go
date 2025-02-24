package repository

import "amartha-go-struct/entity"

type DiscountRepository interface {
	Create(discount entity.DiscountEntity) (entity.DiscountEntity, error)
	FindAll() ([]entity.DiscountEntity, error)
	Update(category entity.DiscountEntity) (entity.DiscountEntity, error)
	Delete(category entity.DiscountEntity) (entity.DiscountEntity, error)
}

type DiscountRepositoryImpl struct {
	discountRepository DiscountRepository
}

func (repo *DiscountRepositoryImpl) Create(discount entity.DiscountEntity) (entity.DiscountEntity, error) {
	return repo.discountRepository.Create(discount)
}

func (repo *DiscountRepositoryImpl) FindAll() ([]entity.DiscountEntity, error) {
	return repo.discountRepository.FindAll()
}

func (repo *DiscountRepositoryImpl) Update(category entity.DiscountEntity) (entity.DiscountEntity, error) {
	return repo.discountRepository.Update(category)
}

func (repo *DiscountRepositoryImpl) Delete(category entity.DiscountEntity) (entity.DiscountEntity, error) {
	return repo.discountRepository.Delete(category)
}

func NewDiscountRepositoryImpl(discountRepository DiscountRepository) *DiscountRepositoryImpl {
	return &DiscountRepositoryImpl{discountRepository: discountRepository}
}
