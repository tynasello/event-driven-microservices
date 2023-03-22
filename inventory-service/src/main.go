package main

import (
	"example.com/inventory-service/src/application/usecase"
	"example.com/inventory-service/src/infra/repository"
	"example.com/inventory-service/src/infra/rest"
	"example.com/inventory-service/src/infra/rest/controllers"
	"example.com/inventory-service/src/infra/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	var dbService service.DbService
	dbService.ConnectToDb()
	dbService.RunDbMigrations()

	// dependency injection
	inventoryRepository := repository.InventoryRepository{Db: dbService.Db}
	createInventoryUseCase := usecase.CreateInventoryUseCase{InventoryRepository: &inventoryRepository}
	inventoryController := controllers.InventoryController{CreateInventoryUseCase: createInventoryUseCase}
	httpServer := rest.HttpServer{InventoryController: inventoryController}

	httpServer.ServeHTTP()
}
