package main

import (
	"fmt"
	"os"
	"io"
)

// Logger interface
type Logger interface {
	Log(message string)
}


type FileLogger struct{
	file io.Writer
}

// ConsoleLogger struct
type ConsoleLogger struct{
	out io.ReadWriter
}

type LogSystem struct{
	logger Logger
}

// LogOption functional option type
type LogOption func(*LogSystem)

func NewLogSystem(options ...LogOption) *LogSystem {
	ls := &LogSystem{}
	for _, option := range options {
		option(ls)
	}
	return ls
}

func WithLogger(l Logger) LogOption {
	return func(ls *LogSystem) {
		ls.logger = l
	}
}

func (f *FileLogger) Log(message string) {
	fmt.Fprintln(f.file, message)
}

// Implementing the Log method for ConsoleLogger
func (c *ConsoleLogger) Log(message string) {
	fmt.Fprintln(c.out, message)
}

func (ls *LogSystem) Log(message string) {
	if ls.logger != nil {
		ls.logger.Log(message)
	}
}

func main() {
	file, _ := os.Create("log.txt")
	defer file.Close()

	fileLogger := &FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))

	logSystem.Log("Hello, world!")
}