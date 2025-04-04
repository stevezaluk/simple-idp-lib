package server

import "github.com/gin-gonic/gin"

/*
Service - An abstraction of a Micro-Service. Responsible for exposing the underlying
gin REST API and providing a connection to MongoDB
*/
type Service struct {
	// Name - The name of the micro-service
	Name string

	// router - The gin router responsible for
	router *gin.Engine

	// database - The MongoDB database that this Service is connected to
	database *Database
}

/*
New - A constructor for the Service object
*/
func New(name string) *Service {
	return &Service{
		Name: name,
	}
}

/*
FromConfig - A wrapper for New that initializes the micro-service from
predefined values in Viper
*/
func FromConfig() *Service {
	return &Service{}
}
