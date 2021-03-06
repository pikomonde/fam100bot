package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/pikomonde/fam100bot/baktu"
	"github.com/pikomonde/fam100bot/fambot"
	cli "github.com/pikomonde/fam100bot/io/client"
	"github.com/pikomonde/fam100bot/tictac"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/ping", ping)

	client := cli.New()

	tictacModule := tictac.New(r, tictac.Option{})
	tictacModule.Register()

	baktuModule := baktu.New(r, baktu.Option{
		Client: client,
	})
	baktuModule.Register()

	fambotModule := fambot.New(r, fambot.Option{
		Client: client,
	})
	fambotModule.Register()

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "7481"
	}
	go r.Run(":" + port)

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
