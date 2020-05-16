package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Goならわかるシステムプログラミング（PDF版）.pdf page 247

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	s := <-sigchan

	switch s {
	case syscall.SIGINT:
		fmt.Println("SIGINT")
	case syscall.SIGTERM:
		fmt.Println("SIGTERM")
	}
}
