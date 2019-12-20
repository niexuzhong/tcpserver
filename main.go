package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"tcpserver/controllers"

	"time"
)

func main() {
	router := gin.Default()
	router.StaticFS("/static", http.Dir("static"))
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.LoadHTMLGlob("views/*")
	router.GET("/index.html", controllers.IndexHandler)
	router.GET("/ws/init", controllers.Initwebsocket)
	//v := router.Group("/")
	//v.GET("/index.html", controllers.IndexHandler)
	srv := &http.Server{
		Addr:         ":9001",
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("shut down server ....")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server shut down:", err)
	}
	log.Println("server exitting")
}
