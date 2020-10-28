package cmd

import (
	"fmt"
	"log"

	"github.com/axard/todo-list/cmd/client/internal/transport"
	"github.com/axard/todo-list/internal/client"
	"github.com/axard/todo-list/internal/client/todos"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/spf13/cobra"
)

const (
	defaultLimit = 20
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Прочитать тудушки",
	RunE: func(cmd *cobra.Command, _ []string) error {
		t := transport.New(host)
		c := client.New(t, strfmt.Default)
		p := todos.NewReadTodosParams()

		p.SetSince(swag.Int64(since))
		p.SetLimit(swag.Int32(limit))

		reading, err := c.Todos.ReadTodos(p)
		if err != nil {
			log.Println(err)
			return nil
		}

		fmt.Printf(
			"OK: %d-%d/%d\n",
			since,
			since+int64(limit),
			*reading.Payload.Total,
		)

		fmt.Printf("    получено: %d\n", len(reading.GetPayload().Items))

		for _, item := range reading.Payload.Items {
			fmt.Println("{")
			fmt.Printf(
				"    ID: %d; Description: '%s'\n",
				item.ID,
				swag.StringValue(item.Description),
			)
			fmt.Println("}")
		}

		return nil
	},
}

var (
	since int64 = 0
	limit int32 = defaultLimit
)

func init() {
	readCmd.Flags().StringVarP(
		&host,
		"host",
		"H",
		defaultHost,
		"Адрес хоста с сервером тудушек",
	)

	readCmd.Flags().Int64VarP(
		&since,
		"since",
		"s",
		0,
		"Начало запрашиваемого диапазона",
	)

	readCmd.Flags().Int32VarP(
		&limit,
		"limit",
		"l",
		defaultLimit,
		"Размер запрашиваемого диапазона",
	)

	todoListClient.AddCommand(readCmd)
}
