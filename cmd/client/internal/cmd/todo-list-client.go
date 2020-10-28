package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var todoListClient = &cobra.Command{
	Use:   "client",
	Short: "Пример клиента по работе с API",
}

func Execute() {
	if err := todoListClient.Execute(); err != nil {
		log.Fatalf("Ошибка выполнения команды: %s", err)
	}
}
