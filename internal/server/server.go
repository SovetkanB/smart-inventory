package server

import (
	"github.com/SovetkanB/smart-inventory/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupServer(productHandler *handler.ProductHandler) *gin.Engine {
	r := gin.Default()

	return r
}
