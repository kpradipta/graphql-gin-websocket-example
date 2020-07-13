package main

import (
	"github.com/graphql-gin-websocket-exampler/mqtt"
	"os"

	opLogger "github.com/op/go-logging"
	"github.com/graphql-gin-websocket-exampler/server/config"
	"github.com/graphql-gin-websocket-exampler/server/logging"
)

var log = logging.MustGetLogger("logs")

func main() {

	SetupLogging()
	err := config.InitConfig()
	if err != nil {
		log.Critical("Error initialize config file")
		os.Exit(1)
	}

	log.Info("","############ INIT MQTT ##############")
	mqtt.Init()
	log.Info("","############ INIT WEB ENGINE ##############")
	config.InitWebEngine()
	log.Info("","############ INIT ROUTER ##############")
	config.InitRouter()
	log.Info("","############ INIT DONE CHEERS!! ##############")


}

func SetupLogging() {

	format := opLogger.MustStringFormatter(
		`%{color} %{time:2006-01-02T15:04:05.999Z07:00} %{shortfile:20.20s} %{shortfunc:10.10s} ▶ %{level:.4s} %{message}%{color:reset}`,
	)
	backend := opLogger.NewLogBackend(os.Stdin, "", 0)
	formatter := opLogger.NewBackendFormatter(backend, format)
	opLogger.SetBackend(formatter)
	log.Info(logging.INTERNAL, "logging initialized")

}
