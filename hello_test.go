package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hey Im Hazim's Go!"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
