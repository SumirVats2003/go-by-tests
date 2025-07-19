package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
  t.Run("adding two numbers", func(t *testing.T) {
    sum := Add(2, 2)
    expect := 4
    
    if sum != expect {
      t.Errorf("expected %d but got %d", expect, sum)
    }
  })
}

func ExampleAdd() {
  sum := Add(1, 5)
  fmt.Println(sum)
  // Output: 6
}
