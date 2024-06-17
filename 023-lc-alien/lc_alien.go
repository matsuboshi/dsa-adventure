package main

import (
	"fmt"
	"sort"
)

func alienOrder(words []string) string {
	adj := make(map[byte]map[byte]struct{})
	for _, word := range words {
		for i := range word {
			adj[word[i]] = make(map[byte]struct{})
		}
	}

	for i:=0 ; i<len(words)-1 ; i++ {
		w1, w2 := words[i], words[i+1]
		minLen := min(len(w1), len(w2))
		if len(w1) > len(w2) && w1[:minLen] == w2[:minLen] {
			return ""
		}

		for j:=0 ; j<minLen ; j++ {
			if w1[j] != w2[j] {
				adj[w1[j]][w2[j]] = struct{}{}
				break
			}
		}
	}

	visit := make(map[byte]bool)
	res := []byte{}

	var dfs func(c byte) bool
	dfs = func(c byte) bool {
		if value, ok := visit[c] ; ok {
			return value
		}

		visit[c] = true

		for nei := range adj[c] {
			if dfs(nei) {
				return true
			}
		}

		visit[c] = false
		res = append(res, c)
		return false
	}

	for c := range adj {
		if dfs(c) {
			return ""
		}
	}
 
	sort.Slice(res, func (i, j int) bool {
		return i > j
	})
	return string(res)
}

func main() {
	words := []string{"wrt", "wrf", "er", "ett", "rftt"}
	fmt.Println(alienOrder(words))
}

