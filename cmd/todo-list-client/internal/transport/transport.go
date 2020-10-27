package transport

import (
	api "github.com/axard/todo-list/internal/client"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/client"
)

const (
	ContentType = "application/io.goswagger.examples.todo-list.v1+json"
)

func New(host string) *client.Runtime {
	rc := client.New(host, api.DefaultBasePath, api.DefaultSchemes)

	rc.Producers[ContentType] = runtime.JSONProducer()
	rc.Consumers[ContentType] = runtime.JSONConsumer()

	return rc
}
