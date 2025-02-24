package repository

import "amartha-go-struct/entity"

type CustomerRepository interface {
	Create(category entity.CustomerEntity) (entity.CustomerEntity, error)
	FindAll() ([]entity.CustomerEntity, error)
	Update(category entity.CustomerEntity) (entity.CustomerEntity, error)
	Delete(category entity.CustomerEntity) (entity.CustomerEntity, error)
}

type CustomerRepositoryImpl struct {
	customerRepository CustomerRepository
}

func (repo *CustomerRepositoryImpl) Create(category entity.CustomerEntity) (entity.CustomerEntity, error) {
	return repo.customerRepository.Create(category)
}

func (repo *CustomerRepositoryImpl) FindAll() ([]entity.CustomerEntity, error) {
	return repo.customerRepository.FindAll()
}

func (repo *CustomerRepositoryImpl) Update(category entity.CustomerEntity) (entity.CustomerEntity, error) {
	return repo.customerRepository.Update(category)
}

func (repo *CustomerRepositoryImpl) Delete(category entity.CustomerEntity) (entity.CustomerEntity, error) {
	return repo.customerRepository.Delete(category)
}

func NewCustomerRepository(customerRepository CustomerRepository) CustomerRepository {
	return &CustomerRepositoryImpl{customerRepository: customerRepository}
}
