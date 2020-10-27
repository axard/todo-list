package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/axard/todo-list/cmd/todo-list-client/internal/transport"
	"github.com/axard/todo-list/internal/client"
	"github.com/axard/todo-list/internal/client/todos"
	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Удалить тудушку",
	RunE: func(cmd *cobra.Command, _ []string) error {
		if id == 0 {
			return errors.New("Не указан id тудушки")
		}

		t := transport.New(host)
		c := client.New(t, strfmt.Default)
		p := todos.NewDeleteOneParams()

		p.SetID(id)

		_, err := c.Todos.DeleteOne(p)
		if err != nil {
			log.Println(err)
			return nil
		}

		fmt.Printf("OK\n")

		return nil
	},
}

var (
	id int64 = 0
)

func init() {
	deleteCmd.Flags().StringVarP(
		&host,
		"host",
		"H",
		defaultHost,
		"Адрес хоста с сервером тудушек",
	)

	deleteCmd.Flags().Int64VarP(
		&id,
		"id",
		"i",
		0,
		"Идентификатор тудушки",
	)

	todoListClient.AddCommand(deleteCmd)
}
