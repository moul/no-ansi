package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/kr/pty"
	"github.com/moul/no-ansi"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	var wg sync.WaitGroup
	if len(os.Args) < 2 {
		// Read from stdin
		noansi.NoAnsiStream(os.Stdin, os.Stdout, &wg)
	} else {
		// Executing a program
		spawn := exec.Command(os.Args[1], os.Args[2:]...)

		// Setup tty
		tty, err := pty.Start(spawn)
		if err != nil {
			log.Fatalln(err)
		}
		defer tty.Close()

		// Setup raw input terminal
		oldState, err := terminal.MakeRaw(0)
		if err == nil {
			defer terminal.Restore(0, oldState)
		}

		// Process stdout routine
		noansi.NoAnsiStream(tty, os.Stdout, &wg)

		// Forward stdin routine
		go func() {
			wg.Add(1)
			defer wg.Done()
			io.Copy(tty, os.Stdin)
		}()

		// Wait for program to finish
		if err := spawn.Wait(); err != nil {
			panic(err)
		}
	}
	wg.Wait()
}
