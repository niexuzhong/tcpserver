package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func  TestFileDownHandler(t *testing.T) {
	router := gin.Default()
	router.StaticFS("../static", http.Dir("static"))
	router.StaticFile("../favicon.ico", "./favicon.ico")
	router.LoadHTMLGlob("../views/*")
	router.GET("../index.html", IndexHandler)
	router.GET("../ws/init", Initwebsocket)
	log.Println("test begin")
	router.GET("/download",FileDownHandler)
	req, _ := http.NewRequest("GET", "/download", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t,http.StatusOK,w.Code)
	t.Log(w.Body)
}