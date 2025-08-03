package main

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("CI")
	want := "Hello, CI"

	if got != want {
		t.Errorf("Greet() = %q, want %q", got, want)
	}
}
