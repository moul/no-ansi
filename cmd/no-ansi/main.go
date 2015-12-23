package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/moul/no-ansi"
)

func NoAnsiStream(inputStream io.Reader, outputStream io.Writer) <-chan bool {
	finished := make(chan bool, 1)

	go func() {
		scanner := bufio.NewScanner(inputStream)

		for scanner.Scan() {
			line := scanner.Text()

			output, err := noansi.NoAnsiString(line)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(outputStream, output)
		}

		finished <- true
	}()

	return finished
}

func main() {
	if len(os.Args) < 2 {
		// Read from stdin
		<-NoAnsiStream(os.Stdin, os.Stdout)
	} else {
		// Executing a program
		spawn := exec.Command(os.Args[1], os.Args[2:]...)

		// Pipe stdin
		spawn.Stdin = os.Stdin

		// Create reader objects for stdout and stderr
		stdout, err := spawn.StdoutPipe()
		if err != nil {
			panic(err)
		}
		stderr, err := spawn.StderrPipe()
		if err != nil {
			panic(err)
		}

		// Start
		if err := spawn.Start(); err != nil {
			panic(err)
		}

		// Create routines for stdout and stderr
		outFinished := NoAnsiStream(stdout, os.Stdout)
		errFinished := NoAnsiStream(stderr, os.Stderr)

		// Wait for program to finish
		if err := spawn.Wait(); err != nil {
			panic(err)
		}

		wait := 2
		for wait > 0 {
			select {
			case <-outFinished:
				wait--
			case <-errFinished:
				wait--
			}
		}
	}
}
