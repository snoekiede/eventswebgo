package models

import "time"

type WebEvent struct {
	Id          uint `gorm:primaryKey,autoIncrement`
	Title       string
	Description string
	Location    string
	StartDate   *time.Time
	EndDate     *time.Time
}
