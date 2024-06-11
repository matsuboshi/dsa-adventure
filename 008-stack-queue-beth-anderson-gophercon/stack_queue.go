package main

import "fmt"

//////////////////////////////////////////////////
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

///////////////////////////////////////////////////
type LinkedList[T any] struct {
	Head  *Node[T]
	Tail  *Node[T]
	Count int
}

func (l *LinkedList[T]) AddHead(n *Node[T]) {
	temp := l.Head
	l.Head = n
	l.Head.Next = temp
	l.Count++
	if l.Count == 1 {
		l.Tail = l.Head
	}
}

func (l *LinkedList[T]) AddTail(n *Node[T]) {
	if l.Count == 0 {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
	}

	l.Tail = n
	l.Count++
}

func (l *LinkedList[T]) RemoveHead() {
	if l.Count != 0 {
		l.Head = l.Head.Next
		l.Count--
		if l.Count == 0 {
			l.Tail = nil
		}
	}
}

func (l *LinkedList[T]) RemoveTail() {
	if l.Count != 0 {
		if l.Count == 1 {
			l.Head = nil
			l.Tail = nil
		} else {
			current := l.Head
			for current.Next != l.Tail {
				current = current.Next
			}
			current.Next = nil
			l.Tail = current
		}
		l.Count--
	}
}

/////////////////////////////////////////////
type Stack[T any] struct {
	linkedList LinkedList[T]
}

func (s *Stack[T]) Push(o T) {
	node := Node[T]{}
	node.Value = o
	s.linkedList.AddHead(&node)
}

func (s *Stack[T]) Peek() (T, error) {
	var r T
	if s.linkedList.Count == 0 {
		return r, fmt.Errorf("error: empty stack")
	}
	return s.linkedList.Head.Value, nil
}

func (s *Stack[T]) Pop() (T, error) {
	var r T
	if s.linkedList.Count == 0 {
		return r, fmt.Errorf("error: empty stack")
	}
	r, err := s.Peek()
	if err != nil {
		return r, err
	}
	s.linkedList.RemoveHead()

	return r, nil
}

/////////////////////////////////////////////
type Queue[T any] struct {
	linkedList LinkedList[T]
}

func (q *Queue[T]) Enqueue(o T) {
	node := Node[T]{}
	node.Value = o
	q.linkedList.AddTail(&node)
}

func (q *Queue[T]) Dequeue() error {
	if q.linkedList.Count == 0 {
		return fmt.Errorf("error: empty queue")
	}
	q.linkedList.RemoveHead()

	return nil
}

func (q *Queue[T]) Peek() (T, error) {
	var r T
	if q.linkedList.Count == 0 {
		return r, fmt.Errorf("error: empty queue")
	}

	r = q.linkedList.Head.Value
	return r, nil
}


///////////////////////////////////////////
///////////////////////////////////////////
func main() {
	tail := Node[int]{Value: 4}
	scnd := Node[int]{Value: 3, Next: &tail}
	frst := Node[int]{Value: 2, Next: &scnd}
	head := Node[int]{Value: 1, Next: &frst}

	frst.Next = &tail
	scnd.Next = &frst
	head.Next = &scnd

	curr := &head
	println(curr.Value)

	for curr.Next != nil {
		curr = curr.Next
		println(curr.Value)
	}

	// LINKED LIST ///////////////////////////
	ll := &LinkedList[int]{}
	fmt.Println(ll)
	ll.AddHead(&(Node[int]{Value:4}))
	ll.AddHead(&(Node[int]{Value:6}))
	ll.AddHead(&(Node[int]{Value:10}))
	fmt.Println(ll)
	ll.RemoveTail()
	fmt.Println(ll)
	ll.RemoveHead()
	fmt.Println(ll)

	// STACK //////////////////////////
	stack := &Stack[string]{}
	stack.Push("ONE")
	stack.Push("TWO")
	o, _ := stack.Pop() 
	println(o)
	o, _ = stack.Pop() 
	println(o)
	o, err := stack.Pop() 
	println(o, err.Error())
	
	// QUEUE /////////////////////////////
	queue := &Queue[string]{}
	queue.Enqueue("FIRST IN")
	queue.Enqueue("SECOND IN")
	o, _ = queue.Peek()
	println(o)
	queue.Dequeue()
	o, _ = queue.Peek()
	println(o)
	queue.Dequeue()
	o, err = queue.Peek()
	println(o, err.Error())
}
