package baktu

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Module contains all config needed
type Module struct {
	Router *gin.Engine
	Prefix string
	//DB *database.Store
}

// Option contains all config needed
type Option struct {
	//Prefix string
}

// New set config to a module
func New(router *gin.Engine, opt Option) *Module {
	return &Module{
		Router: router,
		//Prefix: opt.Prefix,
	}
}

// Register the endpoints
func (m *Module) Register() {
	m.Router.GET("/baktu/ping", ping)
	m.Router.GET("/baktu/direct", ping)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "baktu pong",
	})
}
