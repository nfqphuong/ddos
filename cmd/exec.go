package cmd

import (
	"fmt"
	"os"
)

// Executor - helps execute the command
type Executor struct {
	config *CommandConfig
}

// Printf - prints out message with format
func (e Executor) Printf(f string, a ...interface{}) {
	fmt.Fprintf(e.config.OutputStream, f, a...)
}

// Println - prints out message and line break
func (e Executor) Println(a ...interface{}) {
	fmt.Fprintln(e.config.OutputStream, a...)
}

// Terminate - end commend process
func (e Executor) Terminate() {
	os.Exit(1)
}

// NewExecutor - create new executor
func NewExecutor() *Executor {
	return &Executor{
		&CommandConfig{os.Stderr},
	}
}
