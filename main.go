/*
CLI tool для подсчёта слов в инпуте из stdin.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

/*
Считает количество слов или строк в переданном ридере.
	- countLines: считать строки, если true
*/
func count(r io.Reader, countLines bool) int {
	scanner := bufio.NewScanner(r)

	// Default behavior for Split() is split by lines
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	linesFlag := flag.Bool("l", false, "Count lines")
	flag.Parse()

	result := count(os.Stdin, *linesFlag)
	fmt.Println(result)
}