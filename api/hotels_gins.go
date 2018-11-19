package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1/hotels")
	{
		v1.POST("/", createHotel)
		v1.GET("/", getAllHotels)
		v1.GET("/:id", getHotel)
		v1.PUT("/:id", updateHotel)
		v1.DELETE("/:id", deleteHotel)
	}
	router.Run()
}

type Hotel struct {
	Id          string `json:"id" binding:"required"`
	DisplayName string `json:"display_name" `
	StarRating  int    `json:"star_rating" `
	NoRooms     int    `json:"no_rooms" `
	Links       []Link `json:"links"`
}

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
	Type string `json:"type"`
}

var (
	repository map[string]*Hotel
)

func init() {
	repository = make(map[string]*Hotel)
}

func (h *Hotel) generateHateosLinks(url string) {
	// Book url
	postLink := Link{
		Href: url + "book",
		Rel:  "book",
		Type: "POST",
	}

	h.Links = append(h.Links, postLink)
}

func createHotel(c *gin.Context) {
	var hotel Hotel
	if err := c.ShouldBindJSON(&hotel); err == nil {
		// add HATEOS links
		hotel.generateHateosLinks(c.Request.URL.String())
		// add hotel to repository
		repository[hotel.Id] = &hotel

		//return OK
		c.JSON(http.StatusAccepted, gin.H{"status": "created"})
	} else {
		// some params not correct
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func getHotel(c *gin.Context) {
	// get ID from path param
	hotelId := c.Param("id")

	// get hotel object from repository
	hotel, found := repository[hotelId]
	fmt.Println(hotel, found, hotelId)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"status": "hotel with id not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": hotel})
	}

}

func updateHotel(c *gin.Context) {
	// get hotel object from repository
	hotelId := c.Param("id")
	hotel, found := repository[hotelId]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"status": "hotel with id not found"})
	} else {
		//update
		if err := c.ShouldBindJSON(&hotel); err == nil {
			repository[hotel.Id] = hotel

			//return OK
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		} else {
			// some params not correct
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	}

}

func getAllHotels(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"result": repository})
}

func deleteHotel(c *gin.Context) {
	hotelId := c.Param("id")
	_, found := repository[hotelId]
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"status": "hotel with id not found"})
	} else {
		delete(repository, hotelId)
		//return OK
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}
