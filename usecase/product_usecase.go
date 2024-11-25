package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repository repository.ProductRepository) ProductUsecase {
	return ProductUsecase{repository: repository}
}

func (usecase *ProductUsecase) GetProducts() ([]model.Product, error) {
	return usecase.repository.GetProducts()
}

func (usecase *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := usecase.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.Id = productId
	return product, nil
}

func (usecase *ProductUsecase) GetProductById(id int) (*model.Product, error) {
	product, err := usecase.repository.GetProductById(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
