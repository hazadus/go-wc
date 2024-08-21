/*
CLI tool для подсчёта строк, слов или байтов в инпуте из stdin.
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
	- countBytes: считать байты, если true
*/
func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	// Default behavior for Split() is split by lines
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	linesFlag := flag.Bool("l", false, "Count lines")
	bytesFlag := flag.Bool("b", false, "Count lines")
	flag.Parse()

	result := count(os.Stdin, *linesFlag, *bytesFlag)
	fmt.Println(result)
}