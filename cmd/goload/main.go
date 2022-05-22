package main

import (
	"goload/internal"
)

func main() {
	config := internal.Bootstrap()

	executor := internal.NewExecutor(config)
	executor.Execute()
}
