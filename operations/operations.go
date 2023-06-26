package operations

import (
	"eventsWeb/db"
	"eventsWeb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	var webEvent models.WebEvent
	go func() {
		c.BindJSON(&webEvent)
		db, err := db.GetConnection()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		db.Connection.Create(&webEvent)
		c.JSON(http.StatusCreated, webEvent)
	}()
}

func ListEvents(c *gin.Context) {
	result := make(chan []models.WebEvent)
	go func(context *gin.Context) {
		var webEvents []models.WebEvent

		db, err := db.GetConnection()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}
		db.Connection.Find(&webEvents)
		result <- webEvents
	}(c.Copy())
	c.JSON(http.StatusOK, <-result)
}

func FindEvent(c *gin.Context) {
	var webEvent models.WebEvent
	go func() {
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
	}()
}

func DeleteEvent(c *gin.Context) {
	go func() {
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
	}()
}
