package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataStore interface {
	GetData() string
	SaveData(data string)
}

type InMemoryDataStore struct {
	data string
}

func (imd *InMemoryDataStore) GetData() string {
	return imd.data
}

func (imd *InMemoryDataStore) SaveData(data string) {
	imd.data = data

}

func main() {
	dataStore := &InMemoryDataStore{}

	router := gin.Default()

	chnl := make(chan string)

	router.GET("/data", func(c *gin.Context) {
		go func() {
			data := dataStore.GetData()

			fmt.Println("Data extracted is", data)

			fmt.Println("Data received from Channel is", <-chnl)

		}()

		c.JSON(http.StatusOK, gin.H{
			"message": "Data request received",
		})
	})

	router.POST("/data", func(c *gin.Context) {
		var requestData struct {
			Data string `json:"data"`
		}
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		go func() {
			dataStore.SaveData(requestData.Data)
			fmt.Println("Data saved:", requestData.Data)

			chnl <- requestData.Data

			fmt.Println("Data also received in Channel")
		}()

		c.JSON(http.StatusOK, gin.H{
			"message": "Data received and processing started",
		})
	})
	//close(chnl)
	log.Fatal(router.Run(":8080"))
}

//  curl -X GET http://localhost:8080/data
//  curl -X POST -H "Content-Type: application/json" -d '{"data":"example"}' http://localhost:8080/data
