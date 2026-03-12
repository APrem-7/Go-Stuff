package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)
type car struct{
	miles int 
	electric bool
	petrol bool
}

func whatCar(){
	var thisCar car = car{24,true,false}

	if thisCar.electric{
		fmt.Println("yes")
	}
	else{
		fmt.println(thisCar.petrol)
	}
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}


