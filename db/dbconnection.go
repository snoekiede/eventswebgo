package db

import (
	"eventsWeb/models"
	"gorm.io/gorm"
	"sync"
)

type DbConnection struct {
	Connection *gorm.DB
}

var (
	connectionInstance *DbConnection
	once               sync.Once
	globalError        error
)

func GetConnection() (*DbConnection, error) {
	globalError = nil

	once.Do(func() {
		dbConnection, dbErr := InitializeDatabaseConnection()
		if dbErr != nil {
			globalError = dbErr
			connectionInstance = nil
			return
		}
		migrateError := dbConnection.AutoMigrate(&models.WebEvent{})
		if migrateError != nil {
			globalError = migrateError
			connectionInstance = nil
			return
		}
		connectionInstance = &DbConnection{
			Connection: dbConnection,
		}

	})
	return connectionInstance, globalError
}
