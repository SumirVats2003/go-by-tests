package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
  t.Run("repeating a character 5 times", func(t *testing.T) {
    got := Repeat("a", 5)
    expect := "aaaaa"

    if got != expect {
      t.Errorf("expected %q and got %q", expect, got)
    }
  })
}

func BenchmarkRepeat(b *testing.B) {
  for b.Loop() {
    Repeat("a", 5)
  }
}

func ExampleRepeat() {
  repeat := Repeat("b", 7)
  fmt.Println(repeat)
  // Output: bbbbbbb
}
