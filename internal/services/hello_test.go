package services

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, world!"
	result := Hello()

	if result != expected {
		t.Errorf("Ожидалось '%s', но получено '%s'", expected, result)
	}
}
