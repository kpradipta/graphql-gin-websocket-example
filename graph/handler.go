package graph

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/warung-pintar/server/logging"
)

var log = logging.MustGetLogger("messaging")

var MessageGraphQL = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Message",
		Fields: graphql.Fields{
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func MessageHandler() gin.HandlerFunc{
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"list_message": &graphql.Field{
				Type:    graphql.NewList(MessageGraphQL),
				Resolve: GetMessageResolver,
			},
		},
	})
	// 2

	mutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"SendMessage": &graphql.Field{
				Type: MessageGraphQL,
				Args: graphql.FieldConfigArgument{
					"text": &graphql.ArgumentConfig{
						Description: "text message",
						Type:        graphql.NewNonNull(graphql.String),
					},
				},
				// 3
				Resolve: SendMessageResolver,
			},
		},
	})
	// 4
	Schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: query,
		Mutation: mutation,
	})

	if err != nil {
		// panic if there is an error in schema
		panic(err)
	}

	// 5
	h :=handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func RealTime() gin.HandlerFunc{
	messaging := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"realtime": &graphql.Field{
				Type:    graphql.NewList(graphql.String),
				Resolve: GetMessageResolver,
			},
		},
	})

	// 4
	Schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: messaging,
	})

	if err != nil {
		// panic if there is an error in schema
		panic(err)
	}

	// 5
	h :=handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
