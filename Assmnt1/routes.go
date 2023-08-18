package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/Add", PutInitialEntry)
	router.GET("/Blocks", GetAllBlocks)
	router.GET("/Block:id", GetBlockById)
	router.GET("/AllEntries", PrintDB)
}