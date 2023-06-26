package main

import (
	"eventsWeb/operations"
	"fmt"

	"github.com/gin-gonic/gin"
)

const version = "v1"
const groupName = "api"

func main() {
	//now we can start serving

	router := gin.Default()
	group := router.Group(fmt.Sprintf("/%s/%s", groupName, version))
	{
		group.GET("/events", operations.ListEvents)
		group.GET("/event/:id", operations.FindEvent)
		group.POST("/create", operations.CreateEvent)
		group.DELETE("/delete/:id", operations.DeleteEvent)
	}

	router.Run()
}
