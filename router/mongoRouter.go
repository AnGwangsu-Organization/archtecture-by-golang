package router

import (
	"github.com/gin-gonic/gin"
)

type MongoRouter struct {
	router *Router
}

func NewMongoRouter(router *Router) {
	m := &MongoRouter{
		router: router,
	}

	m.router.GET("/health", m.health)

}

func (m *MongoRouter) health(c *gin.Context) {
	if !c.Writer.Written() {
		c.JSON(200, "test")
	}
}
