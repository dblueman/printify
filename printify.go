package printify

import (
   "fmt"
   "unicode"
)

func String(in string) string {
   var out string

   for _, r := range in {
      if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSymbol(r) || unicode.IsMark(r) ||r == ' ' {
         out += string(r)
         continue
      }

      out += fmt.Sprintf("<%02X>", r)
   }

   return out
}

