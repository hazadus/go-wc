package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := bytes.NewBufferString("w1 w2 w3 w4 w5\n\n")
	exp := 5

	res := count(input, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

func TestCountLines(t *testing.T) {
	input := bytes.NewBufferString("w1\nw2 w3\nw4 w5")
	exp := 3

	res := count(input, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}