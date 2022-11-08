package httpserver

import "github.com/gin-gonic/gin"

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func NewRouter(r *gin.Engine) *Router {
	return &Router{r: r}
}

type Router struct {
	r *gin.Engine
}

func (receiver *Router) GetEngine() *gin.Engine {
	return receiver.r
}

func (receiver *Router) AddGetHandler(route string, handler func(c *gin.Context)) {
	receiver.r.GET(route, handler)
}

func (receiver *Router) AddPostHandler(route string, handler func(c *gin.Context)) {
	receiver.r.POST(route, handler)
}

func AddCors(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	c.Header("Access-Control-Allow-Origin", origin)
}
