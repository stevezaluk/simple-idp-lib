package server

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"strconv"
	"time"
)

/*
Database An abstraction of an active mongodb database connection. The same connection is re-used across
all SDK operations to ensure that we don't exceed the connection pool limit
*/
type Database struct {
	// clientOptions - An options structure built when the constructor is called
	options *options.ClientOptions

	// defaultDatabase - The database the client will connect to by default
	defaultDatabase string

	// client - The MongoDB client that establishes our connection with the Database
	client *mongo.Client

	// database - The MongoDB database that controls our interaction with the data
	database *mongo.Database
}

/*
NewDatabase - A constructor for the database object. This exists in the event the caller
wants to create a new database object without using Viper
*/
func NewDatabase(hostname string, port int, defaultDatabase string) *Database {
	hosts := hostname + ":" + strconv.Itoa(port)

	clientOptions := options.Client().
		SetHosts([]string{hosts}).
		SetDirect(true).
		SetServerSelectionTimeout(30 * time.Second).
		SetTimeout(30 * time.Second)

	return &Database{
		options:         clientOptions,
		defaultDatabase: defaultDatabase,
	}
}
