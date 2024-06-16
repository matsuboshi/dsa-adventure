package main

import (
	"strings"
	"testing"
	"unicode"
)

 
func BenchmarkToLower (b *testing.B) {
	s := "HELLO WORLD"
	
	for i:=0 ; i<b.N ; i++ {
		s = strings.ToLower(s)
	}
} // this one is better


func BenchmarkUnicode (b *testing.B) {
	s := "HELLO WORLD"
	
	for i:=0 ; i<b.N ; i++ {
		s = strings.Map(unicode.ToLower, s)
	}
} // this one takes three times more



 


