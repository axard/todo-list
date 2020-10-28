package main

import (
	"fmt"

	"github.com/axard/todo-list/cmd/client/internal/cmd"
	"github.com/axard/todo-list/pkg/version"
)

func main() {
	fmt.Printf("Version: %s\n", version.Version)
	cmd.Execute()
}
