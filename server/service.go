package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
)

// HandlerFunc - A wrapper for your gin middleware
type HandlerFunc func(service *Service) func(c *gin.Context)

/*
Service - An abstraction of a Micro-Service. Responsible for exposing the underlying
gin REST API and providing a connection to MongoDB
*/
type Service struct {
	// Name - The name of the Service
	Name string

	// Port - The port to expose the API on. Defaults to 8080
	Port int

	// router - The gin router responsible for
	router *gin.Engine

	// database - The MongoDB database that this Service is connected to
	database *Database
}

/*
New - A constructor for the Service object
*/
func New(name string, database *Database) *Service {
	router := gin.New()
	router.Use(gin.Recovery())

	return &Service{
		Name:     name,
		Port:     8080,
		router:   router,
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

/*
Database - Getter function for returning a pointer to the database
*/
func (service *Service) Database() *Database {
	return service.database
}

/*
RegisterEndpoint - Wraps the gin function gin.Engine.Handle and registers a new endpoint with the
router. The 'handler' parameter should be the logic of your endpoint using the HandlerFunc type
*/
func (service *Service) RegisterEndpoint(method string, endpoint string, handler HandlerFunc) {
	service.router.Handle(method, endpoint, handler(service))
}

/*
Run - Start the Service and expose the API to the port defined in Service.Port. Logic
in here should eventually be deprecated so it can gracefully shutdown. Gin doesn't
provide a way natively within its framework to gracefully stop accepting connections
*/
func (service *Service) Run() error {
	err := service.router.Run("0.0.0.0:" + strconv.Itoa(service.Port))
	if err != nil {
		return err
	}

	return nil
}

/*
Shutdown - Gracefully disconnects from the database. This will be improved eventually
as realistically the API should stop accepting new requests before gracefully disconnecting
from MongoDB
*/
func (service *Service) Shutdown() error {
	err := service.database.Client().Disconnect(context.Background())
	if err != nil {
		return err
	}

	return nil
}
