package main

import (
	"ascenda_assessment/logger"
	"ascenda_assessment/services/hotel"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/v1/hotels", func(c *gin.Context) {
		hotelIDs := c.QueryArray("hotel_ids[]")
		destination := c.Query("destination")

		logger.Info(fmt.Sprintf("Get hotels request, hotel_ids: %s, destination: %s", hotelIDs, destination))

		hotels, err := hotel.GetHotels(destination, hotelIDs)

		if err != nil {
			c.JSON(500, err)
		}

		c.JSON(200, hotels)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
