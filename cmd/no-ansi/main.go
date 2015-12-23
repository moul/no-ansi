package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/moul/no-ansi"
)

func main() {
	// Read from stdin
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		output, err := noansi.NoAnsiString(line)
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
	}
}
