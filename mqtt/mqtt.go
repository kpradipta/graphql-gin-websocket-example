package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/graphql-gin-websocket-example/server/logging"
	"os"
)

var client mqtt.Client
var log = logging.MustGetLogger("logs")

var iter = 0

func Init() {
	log.Info("", "START INIT MQTT")
	if client == nil {
		opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883")
		opts.SetClientID("mac-go")

		client = mqtt.NewClient(opts)

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			fmt.Println(token.Error())
		}
	}

}

func Publish(text string) {
	topic := fmt.Sprintf("realtime/%d", iter)
	client.Publish(topic, 0, false, text)
}

func Subscribe() {
	f := func(client mqtt.Client, msg mqtt.Message) {

	}
	if tokens := client.Subscribe("realtime/#", 0, f); tokens.Wait() && tokens.Error() != nil {
		fmt.Println(tokens.Error())
		os.Exit(1)
	}
}
