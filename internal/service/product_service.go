package service

import "github.com/SovetkanB/smart-inventory/internal/repository"

type ProductService struct {
	repo *repository.ProductRepo
}

func NewProductService(productRepo *repository.ProductRepo) *ProductService {
	return &ProductService{repo: productRepo}
}
