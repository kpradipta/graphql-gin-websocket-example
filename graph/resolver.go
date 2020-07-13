package graph
/*
Resolver class to handle business logic
*/
import (
	"github.com/graphql-go/graphql"
	"github.com/warung-pintar/mqtt"
)

var messages []Message

type Message struct {
	Text string
	Status  string
}


func SendMessageResolver(p graphql.ResolveParams) (interface{}, error) {
	text := p.Args["text"].(string)
	response := Message{
		Text: text,
		Status: "Sent",
	}

	messages = append(messages, response)
	//mqtt.Publish(text)
	publish(text)
	return response, nil
}

func GetMessageResolver(p graphql.ResolveParams) (interface{}, error) {
	return messages, nil
}

func RealtimeResolver(p graphql.ResolveParams) (interface{}, error) {
	mqtt.Subscribe()
	return messages, nil
}
