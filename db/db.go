package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabaseConnection() (*gorm.DB, error) {
	dsn := ConstructDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in InitializeDatabaseConnection", r)
		}
	}()
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
	password := getEnvironmentVariableWithDefault("password", "<Your password>")
	dbname := getEnvironmentVariableWithDefault("dbname", "webevents")
	port := getEnvironmentVariableWithDefault("port", "5432")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

	return dsn
}
