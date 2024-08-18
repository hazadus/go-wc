/*
CLI tool для подсчёта слов в инпуте из stdin.
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Считает количество слов в переданном ридере.
func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	result := count(os.Stdin)
	fmt.Println(result)
}