package main

import (
	"example.com/inventory-service/src/internal/config"
	"example.com/inventory-service/src/internal/rest"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDb()
	config.RunDbMigrations()
}

func main() {
	rest.Init()
}
