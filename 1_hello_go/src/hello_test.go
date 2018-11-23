package main

import "testing"

func TestHello(t *testing.T) {

	// Helper function
	assertStr := func(t *testing.T, expected, actual string) {
		// To resolve tracking down problem
		t.Helper()
		if actual != expected {
			t.Errorf("TOBE : '%s', ASIS: '%s'", expected, actual)
		}
	}

	// Test Subset 1
	// Requirement : Specify greeting
	t.Run("When call Hello() with args, Should display greeting and args", func(t *testing.T) {
		// Dynamic var definition
		actual := Hello("World", "English")
		expected := "Hello, World"
		assertStr(t, expected, actual)
	})

	// Test Subset 2
	// Requirement : Specify default greeting
	t.Run("When call Hello() with no args, Should display greeting like just Hello", func(t *testing.T) {
		actual := Hello("", "English")
		expected := "Hello, Stranger"

		assertStr(t, expected, actual)
	})

	// Test Subset 3
	// Requirement : Support second param, specifying the lang of the greeting.
	t.Run("When call Hello() with name and specify the lang for Spanish, Should greet Hola", func(t *testing.T) {
		actual := Hello("Elodie", "Spanish")
		expected := "Hola, Elodie"

		assertStr(t, expected, actual)
	})

	// Test Subset 3-1
	t.Run("When call Hello() with name and specify the lang for French, Should greet Bonjour", func(t *testing.T) {
		actual := Hello("Beth", "French")
		expected := "Bonjour, Beth"

		assertStr(t, expected, actual)
	})

}
