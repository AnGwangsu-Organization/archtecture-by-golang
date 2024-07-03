package router

import "github.com/gin-gonic/gin"

func (r *Router) GET(path string, handlers gin.HandlerFunc) gin.IRoutes {
	return r.engin.GET(path, handlers)
}
func (r *Router) POST(path string, handlers gin.HandlerFunc) gin.IRoutes {
	return r.engin.POST(path, handlers)
}
func (r *Router) PUT(path string, handlers gin.HandlerFunc) gin.IRoutes {
	return r.engin.PUT(path, handlers)
}
func (r *Router) PATCH(path string, handlers gin.HandlerFunc) gin.IRoutes {
	return r.engin.PATCH(path, handlers)
}
func (r *Router) DELETE(path string, handlers gin.HandlerFunc) gin.IRoutes {
	return r.engin.DELETE(path, handlers)
}
