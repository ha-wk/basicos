package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

// GET request to retrieve a value from LevelDB based on a key
func GetById(c *gin.Context) {
	key := c.Query("key")

	value, err := db.Get([]byte(key), &opt.ReadOptions{false, opt.StrictAll})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": string(value)})
}

// PUT request to store a value in LevelDB with a given key
func AddNew(c *gin.Context) {
	var data struct {
		Key   string `json:"key" binding:"required"`
		Value string `json:"value" binding:"required"`
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Put([]byte(data.Key), []byte(data.Value), nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Value stored successfully"})
}

// GETALL request to retrieve all key-value pairs from LevelDB
func GetAll(c *gin.Context) {
	iter := db.NewIterator(nil, nil)
	defer iter.Release()

	result := make(map[string]string)
	for iter.Next() {
		key := string(iter.Key())
		value := string(iter.Value())
		result[key] = value
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}
