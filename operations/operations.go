package operations

import (
	"eventsWeb/db"
	"eventsWeb/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEvent(c *gin.Context) {
	result := make(chan models.OperationResult[models.WebEvent])
	var webEvent models.WebEvent
	dbConnection, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if dbConnection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No database connection"})
		return
	}
	go func(context *gin.Context) {
		conversionError := c.BindJSON(&webEvent)
		if conversionError != nil {
			result <- models.ConstructWithError[models.WebEvent](conversionError)
		} else {

			dbResult := dbConnection.Connection.Create(&webEvent)
			if dbResult.Error != nil {
				result <- models.ConstructWithError[models.WebEvent](dbResult.Error)
			} else {
				result <- models.ConstructWithoutError[models.WebEvent](&webEvent)
			}
		}

	}(c.Copy())
	finalResult := <-result
	if finalResult.Result != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": finalResult.Result})
	} else {
		c.JSON(http.StatusOK, finalResult.Value)
	}
}

func ListEvents(c *gin.Context) {
	result := make(chan []models.WebEvent)

	dbConnection, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if dbConnection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No database connection"})
		return
	}

	go func(context *gin.Context) {
		var webEvents []models.WebEvent

		dbConnection.Connection.Find(&webEvents)
		result <- webEvents

	}(c.Copy())
	c.JSON(http.StatusOK, <-result)

}

func FindEvent(c *gin.Context) {
	result := make(chan models.OperationResult[models.WebEvent])
	var webEvent models.WebEvent
	dbConnection, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if dbConnection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No database connection"})
		return
	}
	go func(context *gin.Context) {
		id := c.Param("id")

		dbResult := dbConnection.Connection.First(&webEvent, id)
		if dbResult.Error != nil {
			result <- models.ConstructWithError[models.WebEvent](dbResult.Error)
		} else {
			result <- models.ConstructWithoutError[models.WebEvent](&webEvent)
		}
	}(c.Copy())
	finalResult := <-result
	if finalResult.Result != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": finalResult.Result})
	} else {
		c.JSON(http.StatusOK, finalResult.Value)
	}
}

func DeleteEvent(c *gin.Context) {
	result := make(chan gin.H)
	dbConnection, err := db.GetConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if dbConnection == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No database connection"})
		return
	}
	go func(context *gin.Context) {
		id := c.Param("id")

		dbResult := dbConnection.Connection.Delete(&models.WebEvent{}, id)
		if dbResult.Error != nil {
			result <- gin.H{"message": dbResult.Error}
		} else {
			result <- gin.H{
				"message": "Deleted",
			}
		}
	}(c.Copy())
	c.JSON(http.StatusOK, <-result)
}
