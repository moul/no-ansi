package noansi

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sync"
)

func NoAnsiString(input string) (string, error) {
	r := regexp.MustCompile("\x1b\\[[>0-9;]*[mc]")
	return r.ReplaceAllString(input, ""), nil
}

func NoAnsiStream(inputStream io.Reader, outputStream io.Writer, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(inputStream)

		for scanner.Scan() {
			line := scanner.Text()

			output, err := NoAnsiString(line)
			if err != nil {
				panic(err)
			}
			fmt.Fprintln(outputStream, output)
		}

	}()
}
