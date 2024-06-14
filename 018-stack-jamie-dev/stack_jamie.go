package main
import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items,i)
}

func (s *Stack) Pop() int{
	last := len(s.items)-1
	removed := s.items[last]
	s.items = s.items[:last]
	return removed
}

func main() {
	myStack := &Stack{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	myStack.Push(400)
	fmt.Println(myStack)
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack)
}

