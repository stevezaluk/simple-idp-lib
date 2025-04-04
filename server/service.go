package server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

/*
Service - An abstraction of a Micro-Service. Responsible for exposing the underlying
gin REST API and providing a connection to MongoDB
*/
type Service struct {
	// Name - The name of the Service
	Name string

	// router - The gin router responsible for
	router *gin.Engine

	// database - The MongoDB database that this Service is connected to
	database *Database
}

/*
New - A constructor for the Service object
*/
func New(name string, database *Database) *Service {
	return &Service{
		Name:     name,
		database: database,
	}
}

/*
FromConfig - A wrapper around New. Constructs a new Service
struct using values provided by Viper
*/
func FromConfig() *Service {
	return New(
		viper.GetString("name"),
		NewDatabaseFromConfig(),
	)
}
