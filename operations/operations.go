package operations

import (
	"eventsWeb/db"
	"eventsWeb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	result := make(chan *models.WebEvent)
	var webEvent models.WebEvent
	db, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	go func(context *gin.Context) {
		c.BindJSON(&webEvent)

		dbResult := db.Connection.Create(&webEvent)
		if dbResult.Error != nil {
			result <- nil
		} else {
			result <- &webEvent
		}

	}(c.Copy())
	finalResult := <-result
	if finalResult == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create event"})
	} else {
		c.JSON(http.StatusOK, finalResult)
	}
}

func ListEvents(c *gin.Context) {
	result := make(chan []models.WebEvent)

	db, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	go func(context *gin.Context) {
		var webEvents []models.WebEvent

		db.Connection.Find(&webEvents)
		result <- webEvents

	}(c.Copy())
	c.JSON(http.StatusOK, <-result)

}

func FindEvent(c *gin.Context) {
	result := make(chan *models.WebEvent)
	var webEvent models.WebEvent
	db, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	go func(context *gin.Context) {
		id := c.Param("id")

		dbResult := db.Connection.First(&webEvent, id)
		if dbResult.Error != nil {
			result <- nil
		} else {
			result <- &webEvent
		}
	}(c.Copy())
	var finalResult *models.WebEvent = <-result
	if finalResult == nil {
		c.JSON(http.StatusNotFound, gin.H{})
	} else {
		c.JSON(http.StatusOK, finalResult)
	}
}

func DeleteEvent(c *gin.Context) {
	result := make(chan gin.H)
	db, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	go func(context *gin.Context) {
		id := c.Param("id")

		dbResult := db.Connection.Delete(&models.WebEvent{}, id)
		if dbResult.Error != nil {
			result <- gin.H{"message": "Failed to delete event"}
		} else {
			result <- gin.H{
				"message": "Deleted",
			}
		}
	}(c.Copy())
	c.JSON(http.StatusOK, <-result)
}
