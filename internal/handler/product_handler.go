package handler

import "github.com/SovetkanB/smart-inventory/internal/service"

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{service: productService}
}
