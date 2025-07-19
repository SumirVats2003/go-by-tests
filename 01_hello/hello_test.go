package main

import "testing"

func TestHello(t *testing.T) {
  t.Run("saying hello to people", func(t *testing.T) {
    got := Hello("Evan", "")
    expect := "Hello Evan"
    assertCorrectMessage(t, got, expect)
  })

  t.Run("when an empty string is passed, say Hello World", func(t *testing.T) {
    got := Hello("", "")
    expect := "Hello World"
    assertCorrectMessage(t, got, expect)
  })

  t.Run("greeting in spanish", func(t *testing.T) {
    got := Hello("Dave", "Spanish")
    expect := "Hola Dave"
    assertCorrectMessage(t, got, expect)
  })

  t.Run("greeting in french", func(t *testing.T) {
    got := Hello("Dave", "French")
    expect := "Bonjour Dave"
    assertCorrectMessage(t, got, expect)
  })
}

func assertCorrectMessage(t testing.TB, got, expect string) {
  t.Helper()
  if got != expect {
    t.Errorf("got %q and expected %q", got, expect)
  }
}
