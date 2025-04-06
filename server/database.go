package server

import (
	"context"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log/slog"
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

	// defaultCollection - The collection that the client will use by default
	defaultCollection string

	// client - The MongoDB client that establishes our connection with the Database
	client *mongo.Client

	// database - The MongoDB database that controls our interaction with the data
	database *mongo.Database

	// collection - The MongoDB database that controls our interaction with the collection
	collection *mongo.Collection
}

/*
NewDatabase - A constructor for the database object. This exists in the event the caller
wants to create a new database object without using Viper
*/
func NewDatabase(hostname string, port int, defaultDatabase string, defaultCollection string) *Database {
	hosts := hostname + ":" + strconv.Itoa(port)

	clientOptions := options.Client().
		SetHosts([]string{hosts}).
		SetDirect(true).
		SetServerSelectionTimeout(30 * time.Second).
		SetTimeout(30 * time.Second)

	return &Database{
		options:           clientOptions,
		defaultDatabase:   defaultDatabase,
		defaultCollection: defaultCollection,
	}
}

/*
NewDatabaseFromConfig - A wrapper for NewDatabase. Constructs a new database from
configuration values passed in Viper
*/
func NewDatabaseFromConfig() *Database {
	database := NewDatabase(
		viper.GetString("mongo.hostname"),
		viper.GetInt("mongo.port"),
		viper.GetString("mongo.default_database"),
		viper.GetString("mongo.default_collection"),
	)

	database.SetSCRAMAuthentication(
		viper.GetString("mongo.username"),
		viper.GetString("mongo.password"),
	)

	database.Connect()

	return database
}

/*
Client - Getter function for returning a reference to the MongoDB client
*/
func (database *Database) Client() *mongo.Client {
	return database.client
}

/*
Database - Getter function for returning a pointer to the MongoDB database
*/
func (database *Database) Database() *mongo.Database {
	return database.database
}

/*
Collection - Getter function for returning a pointer to the MongoDB collection
*/
func (database *Database) Collection() *mongo.Collection {
	return database.collection
}

/*
SetSCRAMAuthentication - Set the credentials for the database connection if they are needed
*/
func (database *Database) SetSCRAMAuthentication(username string, password string) {
	credentials := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    "admin",
		Username:      username,
		Password:      password,
	}

	database.options.SetAuth(credentials)
}

/*
Connect to the MongoDB instance defined in the Database object
*/
func (database *Database) Connect() {
	slog.Info("Starting connection to MongoDB")
	client, err := mongo.Connect(database.options)
	if err != nil {
		slog.Error("Failed to connect to database", "err", err)
		panic(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		slog.Error("Failed to ping database after connecting", "err", err)
		panic(err)
	}

	slog.Info("Successfully connected to DB")
	database.client = client
	database.database = database.client.Database(database.defaultDatabase)
	database.collection = database.database.Collection(database.defaultCollection)
}

/*
Find - Fetch a document from MongoDB and decode the results into the reference
passed in the model parameter
*/
func (database *Database) Find(query bson.M, model interface{}) error {
	err := database.collection.FindOne(context.Background(), query).Decode(model)
	if err != nil {
		return err
	}

	return nil
}
