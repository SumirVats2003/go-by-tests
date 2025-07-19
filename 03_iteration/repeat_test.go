package iteration

import "testing"

func TestRepeat(t *testing.T) {
  t.Run("repeating a character 5 times", func(t *testing.T) {
    got := Repeat("a")
    expect := "aaaaa"

    if got != expect {
      t.Errorf("expected %q and got %q", expect, got)
    }
  })
}

func BenchmarkRepeat(b *testing.B) {
  for b.Loop() {
    Repeat("a")
  }
}
