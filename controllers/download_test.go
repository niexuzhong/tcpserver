package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"testing"
	"github.com/Valibern/gin_unit_test"
	"github.com/Valibern/gin_uinit_test/utils"
)

func int()  {
	router := gin.Default()
	router.POST("/donwload",FileDownHandler)
	unitTest.setRounter(router)
}