package main

import "testing"

func TestSprint(t *testing.T) {
	result := Sprint(true, "hello")
	if result != "true, hello" {
		t.Fatal("invalid result.", result)
	}
}
