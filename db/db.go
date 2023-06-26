package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitializeDatabaseConnection() (*gorm.DB, error) {
	dsn := ConstructDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//if err != nil {
	//		panic("Failed to connect to database")
	//}
	return db, err
}

func getEnvironmentVariableWithDefault(key string, defaultValue string) string {
	currentValue := os.Getenv(key)
	if currentValue == "" {
		return defaultValue
	} else {
		return currentValue
	}
}

func ConstructDsn() string {
	host := getEnvironmentVariableWithDefault("host", "localhost")
	user := getEnvironmentVariableWithDefault("user", "postgres")
	password := getEnvironmentVariableWithDefault("password", "Piloten2030")
	dbname := getEnvironmentVariableWithDefault("dbname", "webevents")
	port := getEnvironmentVariableWithDefault("port", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

	return dsn
}
