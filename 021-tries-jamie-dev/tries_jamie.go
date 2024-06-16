package main

import "fmt"

const AlphabetSize = 26

type Node struct {
	children [26]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

func (t *Trie) DFS() {
	var result []string
	var path []rune

	var dfsHelper func(*Node, []rune)
	dfsHelper = func(node *Node, path []rune) {
		if node == nil {
			return
		}
		if node.isEnd {
			result = append(result, string(path))
		}
		for i := 0; i < AlphabetSize; i++ {
			if node.children[i] != nil {
				path = append(path, rune('a'+i))
				dfsHelper(node.children[i], path)
				path = path[:len(path)-1] // Backtrack
			}
		}
	}

	dfsHelper(t.root, path)
	for _, word := range result {
		fmt.Println(word)
	}
}

func main() {
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

	fmt.Println(myTrie.Search("wizard"))
	fmt.Println(myTrie.Search("abcde"))
	fmt.Println(myTrie.Search("mitsuhiko"))

	myTrie.DFS()
}
