package cmd

import (
	"io"
)

// CommandConfig - configration for the whole command process
type CommandConfig struct {
	ThreadNumber  int
	RequestNumber int
	OutputStream  io.Writer
}
