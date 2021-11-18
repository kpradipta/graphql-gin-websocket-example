package socket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/graphql-gin-websocket-example/server/logging"
)

var log = logging.MustGetLogger("socket")

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Subscriber struct {
	ID            string
	Conn          *websocket.Conn
	RequestString string
}

type Request struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	Operation string `json:operation`
}

var Subscribers sync.Map

var subsLength int

// Handler *test code scan
func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		var req Request
		if err := json.Unmarshal(msg, &req); err != nil {
			log.Error("", "error unmarshall json : ", err)
			return
		}

		log.Info("", "REQ : ", req)

		if _, ok := Subscribers.Load(req.Id); !ok {

			Subscribers.Store(conn, &Subscriber{
				ID:            req.Id,
				Conn:          conn,
				RequestString: req.Text,
			})
		}
	}
}
