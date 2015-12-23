package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/kr/pty"
	"github.com/moul/no-ansi"
	"golang.org/x/crypto/ssh/terminal"
)

func NoAnsiStream(inputStream io.Reader, outputStream io.Writer, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(inputStream)

		for scanner.Scan() {
			line := scanner.Text()

			output, err := noansi.NoAnsiString(line)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(outputStream, output)
		}

	}()
}

func main() {
	var wg sync.WaitGroup
	if len(os.Args) < 2 {
		// Read from stdin
		NoAnsiStream(os.Stdin, os.Stdout, &wg)
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
		NoAnsiStream(tty, os.Stdout, &wg)

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
