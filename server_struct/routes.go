package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default() //create gin router

	router.GET("/get", GetById)
	router.PUT("/put", AddNew)
	router.GET("/getall", GetAll)

	return router
}
