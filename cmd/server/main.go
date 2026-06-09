package main

import (
	"log/slog"
	"os"

	"github.com/SovetkanB/smart-inventory/configs"
	"github.com/SovetkanB/smart-inventory/internal/handler"
	"github.com/SovetkanB/smart-inventory/internal/repository"
	"github.com/SovetkanB/smart-inventory/internal/server"
	"github.com/SovetkanB/smart-inventory/internal/service"
	"github.com/SovetkanB/smart-inventory/pkg"
	"github.com/joho/godotenv"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := godotenv.Load(); err != nil {
		slog.Error("No .env file found, using env variables")
		os.Exit(1)
	}

	cfg, err := configs.Load()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	db, err := pkg.Connect(cfg.DB.DSN())
	if err != nil {
		slog.Error("Failed to connect to DB", "error", err)
		os.Exit(1)
	}
	defer db.Close()
	slog.Info("Connected to PostgreSQL")

	productRepo := repository.NewProductRepo(db)
	productService := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	server := server.SetupServer(productHandler)

	if err := server.Run(":8080"); err != nil {
		slog.Error(err.Error())
	}

}
