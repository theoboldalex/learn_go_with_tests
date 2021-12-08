package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Alex")
	expected := "Hello Alex"

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}
