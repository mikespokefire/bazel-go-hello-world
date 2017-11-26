package utils

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello from utils"
	actual := Hello()
	if actual != expected {
		t.Errorf("Hello was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestGoodbye(t *testing.T) {
	expected := "Goodbye from utils"
	actual := Goodbye()
	if actual != expected {
		t.Errorf("Goodbye was incorrect, got: %d, want: %d.", actual, expected)
	}
}
