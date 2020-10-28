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

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Обновить тудушку",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Не указано содержимое тудушки")
		}

		if id == 0 {
			return errors.New("Не указан id тудушки")
		}

		t := transport.New(host)
		c := client.New(t, strfmt.Default)
		p := todos.NewUpdateOneParams()

		p.SetID(id)
		p.SetBody(&restmodels.Item{
			Description: swag.String(args[0]),
			Completed:   done,
		})

		updated, err := c.Todos.UpdateOne(p)
		if err != nil {
			log.Println(err)
			return nil
		}

		fmt.Printf("OK %d\n", updated.Payload.ID)

		return nil
	},
}

var (
	done bool = false
)

func init() {
	updateCmd.Flags().BoolVarP(
		&done,
		"done",
		"d",
		false,
		"Пометить как готовое или не готовое",
	)

	updateCmd.Flags().StringVarP(
		&host,
		"host",
		"H",
		defaultHost,
		"Адрес хоста с сервером тудушек",
	)

	updateCmd.Flags().Int64VarP(
		&id,
		"id",
		"i",
		0,
		"Идентификатор тудушки",
	)

	todoListClient.AddCommand(updateCmd)
}
