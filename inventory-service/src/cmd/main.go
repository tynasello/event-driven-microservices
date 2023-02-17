package main

import (
	"example.com/inventory-service/src/internal/config"
	"example.com/inventory-service/src/internal/rest"
)

func main() {
	config.LoadEnvVariables()
	config.ConnectToDb()
	config.RunDbMigrations()
	rest.ServeHTTP()
}
