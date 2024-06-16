package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key string
	next *bucketNode
}

func (h *HashTable) Print() {
	fmt.Println()
	for i, v := range h.array {
		fmt.Print(i, v)
		cur := v.head
		for cur != nil {
			fmt.Print(" ", cur.key)
			cur = cur.next
		}
		fmt.Println()
	}
	fmt.Println()
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
	
}

func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists!")
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next.key != k {
		if previousNode.next.next == nil {
			println(k, "not found")
			return
		}
		previousNode = previousNode.next
	}
	previousNode.next = previousNode.next.next
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashTable := Init()
	testHashTable.Insert("RANDY")
	testHashTable.Insert("L3")
	testHashTable.Insert("L2")
	testHashTable.Insert("03")
	testHashTable.Insert("04")
	testHashTable.Insert("BOSA")
	testHashTable.Insert("TEMPUREITO")
	testHashTable.Insert("PAGLIACI")
	testHashTable.Insert("LUCSA")
	testHashTable.Insert("LUCAS")

	testHashTable.Insert("EDER")
	testHashTable.Insert("EDER")
	testHashTable.Insert("MITCH")
	testHashTable.Insert("MITCH")

	testHashTable.Print()

	fmt.Println(testHashTable.Search("RANDY"))
	testHashTable.Delete("RANDY")
	testHashTable.Delete("TEMPUREITO")
	testHashTable.Delete("L3")
	fmt.Println(testHashTable.Search("RANDY"))

	testHashTable.Print()
}


