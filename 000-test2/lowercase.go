package main

import (
   "fmt"
   "strings"
   "unicode"
)

func main() {
   s := "HELLO WORLD"
   fmt.Println("Original:", s)

   s1 := strings.ToLower(s)
   fmt.Println("Lowercase:", s1)

   s2 := strings.Map(unicode.ToLower, s)
   fmt.Println("Lowercase:", s2)
}

