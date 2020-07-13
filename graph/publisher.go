package graph
/*
Publisher for real time api
*/
import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/warung-pintar/server/socket"
)

func publish(text string) {
	socket.Subscribers.Range(func(key, value interface{}) bool {
		subscriber, ok := value.(*socket.Subscriber)
		if !ok {
			return true
		}
		if err := subscriber.Conn.WriteMessage(websocket.TextMessage, []byte(text)); err != nil {
			fmt.Println(subscriber.RequestString)
			if err == websocket.ErrCloseSent {
				socket.Subscribers.Delete(key)
				return true
			}
			return true
		}
		return true
	})
}