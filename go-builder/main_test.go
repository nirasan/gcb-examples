package main

import "testing"

func TestHello(t *testing.T) {
	if hello() != "HELLO WORLD!" {
		t.Fatal("invalid string. ", hello())
	}
}
