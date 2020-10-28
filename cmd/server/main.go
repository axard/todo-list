package main

import (
	"log"
	"os"

	"github.com/axard/todo-list/internal/restapi"
	"github.com/axard/todo-list/internal/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
)

func main() {
	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewTodoListAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "A Todo list application"
	parser.LongDescription = "From the todo list tutorial on goswagger.io"

	server.ConfigureFlags()

	// nolint: gocritic
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1

		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}

		os.Exit(code)
	}

	server.ConfigureAPI()

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
