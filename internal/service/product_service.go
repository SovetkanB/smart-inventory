package service

import "github.com/SovetkanB/smart-inventory/internal/repository"

type ProductService struct {
}

func NewProductService(productRepo *repository.ProductRepo) *ProductService {
	return &ProductService{}
}
