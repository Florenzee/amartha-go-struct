package controller

//memanggil repository
import (
	"amartha-go-struct/entity"
	"amartha-go-struct/repository"
)

type CategoryController struct {
	repo repository.CategoryRepository
}

func (controller *CategoryController) Create(category entity.Category) (entity.Category, error) {
	return controller.repo.Create(category)
}

func (controller *CategoryController) Update(category entity.Category) (entity.Category, error) {
	return controller.repo.Update(category)
}
