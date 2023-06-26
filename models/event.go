package models

import "time"

type WebEvent struct {
	Id          uint   `gorm:primaryKey,autoIncrement`
	Title       string `gorm:required`
	Description string
	Location    string
	StartDate   *time.Time
	EndDate     *time.Time
}
