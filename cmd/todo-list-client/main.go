package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/axard/todo-list/internal/client/todos"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	apiclient "github.com/axard/todo-list/internal/client"
	httptransport "github.com/go-openapi/runtime/client"
)

const (
	DefaultLimit = 20
)

var (
	hostFlag = flag.String("host", "localhost", "Host of todo-list service")

	sinceFlag = flag.Int64("since", 0, "Begin of reading")
	limitFlag = flag.Int("limit", DefaultLimit, "Limit of reading")
)

func main() {
	flag.Parse()

	transport := httptransport.New(*hostFlag, apiclient.DefaultBasePath, apiclient.DefaultSchemes)
	transport.Producers["application/io.goswagger.examples.todo-list.v1+json"] = runtime.JSONProducer()
	transport.Consumers["application/io.goswagger.examples.todo-list.v1+json"] = runtime.JSONConsumer()

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	params := todos.NewReadTodosParams()
	params.SetSince(sinceFlag)
	params.SetLimit(swag.Int32(int32(*limitFlag)))

	resp, err := client.Todos.ReadTodos(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total: %d\n", *resp.Payload.Total)

	for _, item := range resp.Payload.Items {
		fmt.Println("{")
		fmt.Printf("    ID: %d; Description: '%s'\n", item.ID, swag.StringValue(item.Description))
		fmt.Println("}")
	}
}
