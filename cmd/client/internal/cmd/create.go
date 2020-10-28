package cmd

import (
	"fmt"
	"log"

	"github.com/axard/todo-list/cmd/client/internal/transport"
	"github.com/axard/todo-list/internal/client"
	"github.com/axard/todo-list/internal/client/todos"
	"github.com/axard/todo-list/internal/restmodels"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	defaultHost = "localhost"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Создать тудушку",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Не указано содержимое тудушки")
		}

		t := transport.New(host)
		c := client.New(t, strfmt.Default)
		p := todos.NewCreateOneParams()

		p.SetBody(&restmodels.Item{
			Description: swag.String(args[0]),
		})

		created, err := c.Todos.CreateOne(p)
		if err != nil {
			log.Println(err)
			return nil
		}

		fmt.Printf("OK: id = %d\n", created.Payload.ID)

		return nil
	},
}

var (
	host string = defaultHost
)

func init() {
	createCmd.Flags().StringVarP(
		&host,
		"host",
		"H",
		defaultHost,
		"Адрес хоста с сервером тудушек",
	)

	todoListClient.AddCommand(createCmd)
}
