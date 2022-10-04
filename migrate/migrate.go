package main

import (
	"Practical_1/initializers"
	models "Practical_1/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.People{})
}