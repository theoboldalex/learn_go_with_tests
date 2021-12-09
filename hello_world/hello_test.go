package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, result string) {
		t.Helper()
		if result != expected {
			t.Errorf("Expected %q but got %q", expected, result)
		} 	
	}

	t.Run("saying hello to people", func(t *testing.T) {
		result := Hello("Alex", "")
		expected := "Hello Alex"

		assertCorrectMessage(t, expected, result)
	})

	t.Run("saying hello world where no name is passed in", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello World"

		assertCorrectMessage(t, expected, result)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		result := Hello("Esteban", "spanish")
		expected := "Hola Esteban"

		assertCorrectMessage(t, expected, result)
	})
	 t.Run("saying hello in french", func(t *testing.T) {
		 result := Hello("Claude", "french")
		 expected := "Bonjour Claude"

		assertCorrectMessage(t, expected, result)
	 })
}
