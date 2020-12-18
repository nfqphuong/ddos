package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nfqphuong/ddos/util"
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
	output := os.Stderr
	var (
		cFlag = flag.Int("c", 1, "total number of threads")
		nFlag = flag.Int("n", 1, "total number of requests")
		oFlag = flag.String("o", "", "write output to a file")
	)
	fmt.Println(*oFlag, *cFlag, *nFlag)
	flag.Parse()
	if len(*oFlag) > 0 {
		fmt.Println("----- wtf")
		f, err := util.GetOutputFile(*oFlag)
		if err != nil {
			log.Fatalln(err)
		}
		output = f
	}

	return &Executor{
		&CommandConfig{*cFlag, *nFlag, output},
	}
}
