package baktu

import (
	"github.com/gin-gonic/gin"
)

// Module contains all config needed
type Module struct {
	router *gin.Engine
	server *server
	//DB *database.Store
}

// Option contains all config needed
type Option struct {
}

// New set config to a module.
func New(router *gin.Engine, opt Option) *Module {
	return &Module{
		router: router,
		server: newServer(),
	}
}

// Register the endpoints.
func (m *Module) Register() {
	m.router.GET("/baktu/ping", m.ping)
	m.router.GET("/baktu/direct", m.direct)
	m.router.GET("/baktu/line/webhook", m.direct)
}
