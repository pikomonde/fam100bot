package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pikomonde/fam100bot/baktu"
	"github.com/pikomonde/fam100bot/tictac"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/ping", ping)

	tictacModule := tictac.New(router, tictac.Option{})
	tictacModule.Register()

	baktuModule := baktu.New(router, baktu.Option{})
	baktuModule.Register()

	go router.Run(":8080")

	term := make(chan os.Signal)
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-term:
		fmt.Println("Exiting gracefully...")
	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "pong",
	})
}
