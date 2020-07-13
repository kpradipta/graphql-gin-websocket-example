package config

import (
	"github.com/gin-gonic/gin"
)

var WebEngine *gin.Engine

func InitWebEngine() {
	serverMode := MustGetString("server.mode")
	debug := MustGetBool(serverMode + ".debug")

	if !debug {
		gin.SetMode(gin.ReleaseMode)
	}
	WebEngine = gin.New()

	if !debug {
		WebEngine.Use(gin.Recovery())
	} else {
		WebEngine.Use(gin.Logger(), gin.Recovery())
	}

}

func GetEngine() *gin.Engine {
	if WebEngine == nil {
		InitWebEngine()
		return WebEngine
	}
	return WebEngine
}
