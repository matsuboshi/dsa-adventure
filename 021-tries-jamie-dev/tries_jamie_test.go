package main

import (
	"fmt"
	"testing"
)

func BenchmarkTries(b *testing.B) {
	myTrie := InitTrie()

	toAdd := []string{
		"abcde",
		"abctr",
		"abcor",
		"abcgg",
		"aragorn",
		"aragon",
		"argon",
		"eragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, v := range toAdd {
		myTrie.Insert(v)
	}

	result := []string{}

	for i:=0 ; i < b.N ; i++ {
		result = myTrie.DFS()
	}

	fmt.Println(result)
}
