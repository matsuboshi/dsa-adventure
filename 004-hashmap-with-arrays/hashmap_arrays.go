package main

import "fmt"

const arraySize = 7

func NewHashMap() *HashMap {
	h := &HashMap{}
	for i := 0; i < arraySize; i++ {
		h.array[i] = &bucket{}
	}
	return h
}

////////////////////////////////
type HashMap struct {
	array [arraySize]*bucket
}

func (h *HashMap) Add(str string) {
	hash := h.hash(str)
	bucket := h.array[hash]
	bucket.add(str)

}

func (h *HashMap) Read() {
	for _, v := range h.array {
		v.read()
	}
}

func (h *HashMap) Contains(str string) {
	hash := h.hash(str)
	bucket := h.array[hash]

	if bucket.find(str) {
		// println(true, str)
		fmt.Printf("%11s %s\n", "CONTAINS", str)
	} else {
		// println(false, str)
		fmt.Printf("%11s %s\n", "NOT CONTAIN", str)
	}
}

func (h *HashMap) Remove(str string) {
	hash := h.hash(str)
	bucket := h.array[hash]

	if bucket.find(str) {
		bucket.delete(str)
		fmt.Printf("%11s %s\n", "DELETED", str)
	} else {
		fmt.Printf("%11s %s\n", "NOT FOUND", str)
	}
}

func (h *HashMap) hash(str string) int {
	divisor := len(h.array)
	sum := 0
	for _, v := range str {
		sum += int(v)
	}
	return sum % divisor
}

// //////////////////////////////
type bucket struct {
	names []string
}

func (b *bucket) read() {
	result := "[ "
	for _, v := range b.names {
		result += v
		result += " "
	}
	result += "]"
	println(result)
}

func (b *bucket) add(str string) {
	b.names = append(b.names, str)
}

func (b *bucket) find(str string) bool {
	for _, v := range b.names {
		if str == v {
			return true
		}
	}
	return false
}

func (b *bucket) delete(str string) {
	for i, v := range b.names {
		if v == str {
			b.names = append(b.names[:i], b.names[i+1:]...)
			return
		}
	}
}

// //////////////////////////////
func main() {
	newHM := NewHashMap()

	// newHM := NewHashMap()
	// for i := 0; i < arraySize; i++ {
	// 	newHM.array[i] = &bucket{}
	// }

	newHM.Add("Mitch")
	newHM.Add("Edinho")
	newHM.Add("Lucas")
	newHM.Add("Tempureito")
	newHM.Read()

	println()
	newHM.Contains("Tempureito")
	newHM.Contains("Mitsuhiko")
	newHM.Contains("Lucas")
	newHM.Contains("Vagabundo")
	println()

	newHM.Remove("Tempureito")
	newHM.Remove("Mitsuhiko")
	newHM.Remove("Edinho")

	println()
	newHM.Read()
}
