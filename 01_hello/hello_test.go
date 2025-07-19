package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("Evan")
  expect := "Hello Evan"

  if got != expect {
    t.Errorf("got %q and expected %q", got, expect)
  }
}
