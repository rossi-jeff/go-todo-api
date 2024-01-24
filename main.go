package main

import (
	"fmt"
	"go-todo-api/initializers"
)

func init() {
	initializers.LoadEnvironment()
	initializers.DatabaseConnect()
}

func main() {
	fmt.Println("main")
}
