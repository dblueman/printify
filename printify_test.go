package printify

import (
   "testing"
)

func TestString(t *testing.T) {
   in := "Hello\nWorld"
   exp := "Hello<0D>World"
   out := String(in)

   if String(in) != out {
      t.Errorf("expected '%s' but got '%s'", exp, out)
   }
}
