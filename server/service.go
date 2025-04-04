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
NewService - A constructor for the Service object
*/
func NewService(name string) *Service {
	return &Service{
		Name: name,
	}
}
