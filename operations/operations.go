package operations

import (
	"eventsWeb/db"
	"eventsWeb/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func TestDbConnection(c *gin.Context) {
	dsn := db.ConstructDsn()
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "DB connection OK",
		})
	}
}

func CreateEvent(c *gin.Context) {
	var webEvent models.WebEvent
	c.BindJSON(&webEvent)
	db, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	db.Connection.Create(&webEvent)
	c.JSON(http.StatusCreated, webEvent)
}

func ListEvents(c *gin.Context) {
	var webEvents []models.WebEvent
	db, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	db.Connection.Find(&webEvents)
	c.JSON(http.StatusOK, webEvents)
}

func FindEvent(c *gin.Context) {
	var webEvent models.WebEvent
	id := c.Param("id")
	db, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	result := db.Connection.First(&webEvent, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(http.StatusOK, webEvent)
	}
}

func DeleteEvent(c *gin.Context) {
	id := c.Param("id")
	db, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	result := db.Connection.Delete(&models.WebEvent{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, nil)
	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"message": "Deleted",
		})
	}
}

func ShowDsn(c *gin.Context) {
	dsn := db.ConstructDsn()
	c.JSON(http.StatusOK, gin.H{
		"dsn": dsn,
	})
}
