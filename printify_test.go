package printify

import (
   "testing"
)

func TestString(t *testing.T) {
   in := "Hello \nWorld>"
   exp := "Hello <0A>World>"
   out := String(in)

   if out != exp {
      t.Errorf("expected '%s' but got '%s'", exp, out)
   }
}
