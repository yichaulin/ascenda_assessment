package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"ascenda_assessment/logger"
	"ascenda_assessment/services/hotel"

	resError "ascenda_assessment/utils/response_error"
)

func main() {
	r := gin.Default()
	r.GET("/api/v1/hotels", func(c *gin.Context) {
		hotelIDs := c.QueryArray("hotel_ids[]")
		destination := c.Query("destination")
		logger.Info(fmt.Sprintf("Get hotels request, hotel_ids: %s, destination: %s", hotelIDs, destination))

		hotels, err := hotel.GetHotels(destination, hotelIDs)
		if err != nil {
			switch err {
			case hotel.InputInvalidError:
				c.AbortWithStatusJSON(http.StatusBadRequest, resError.New(hotel.InputInvalidError))
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, resError.NewInternalServerError(err))
			}
			return
		}

		c.JSON(http.StatusOK, hotels)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
