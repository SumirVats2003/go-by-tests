package iteration

import "strings"

func Repeat(str string) string {
  var repeated strings.Builder
  for range 5 {
    repeated.WriteString(str)
  }
  return repeated.String()
}
