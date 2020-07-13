package config

import (
	"github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/graphql-gin-websocket-exampler/graph"
	"github.com/graphql-gin-websocket-exampler/server/socket"
	"net/http"
)
func InitRouter() {
	port := MustGetString("server.port")
	router := GetEngine()
	router.HTMLRender = gintemplate.Default()


	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.GET("/playground", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "Page file title!!"})
	})

	router.GET("/socket", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "socket.html", gin.H{"title": "Page file title!!"})
	})

	router.GET("/realtime", func(c *gin.Context) {
		socket.WsHandler(c.Writer, c.Request)
	})

	router.POST("/query", graph.MessageHandler() )

	router.Run(":" + port)
}


