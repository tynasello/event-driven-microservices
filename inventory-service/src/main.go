package main

import (
	"example.com/inventory-service/src/application/usecase"
	"example.com/inventory-service/src/infra/messagebroker"
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

	dbService := service.NewDbService()
	dbService.RunDbMigrations()

	// dependency injection
	kafkaBroker := messagebroker.NewKafkaBroker()
	messageBrokerProducerService := service.MessageBrokerProducerService{KafkaBroker: kafkaBroker}

	inventoryRepository := repository.InventoryRepository{Db: dbService.Db}
	createInventoryUseCase := usecase.CreateInventoryUseCase{InventoryRepository: &inventoryRepository}
	reserveInventoryUseCase := usecase.ReserveInventoryUseCase{InventoryRepository: &inventoryRepository, MessageBrokerProducerService: &messageBrokerProducerService}
	orderCancelledUseCase := usecase.OrderCancelledUseCase{InventoryRepository: &inventoryRepository}

	messageBrokerConsumerService := service.MessageBrokerConsumerService{KafkaBroker: kafkaBroker, ReserveInventoryUseCase: reserveInventoryUseCase, OrderCancelledUseCase: orderCancelledUseCase}

	inventoryController := controllers.InventoryController{CreateInventoryUseCase: createInventoryUseCase}
	httpServer := rest.HttpServer{InventoryController: inventoryController}

	go messageBrokerConsumerService.StartConsuming()
	httpServer.ServeHTTP()
}
