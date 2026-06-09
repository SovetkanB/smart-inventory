package handler

import "github.com/SovetkanB/smart-inventory/internal/service"

type ProductHandler struct {
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{}
}
