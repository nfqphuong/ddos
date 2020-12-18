package main

import (
	"fmt"
	"os"

	"github.com/nfqphuong/ddos/cmd"
)

func main() {
	executor := cmd.NewExecutor()
	fmt.Println(os.Args)
	executor.Println("hehe")
}
