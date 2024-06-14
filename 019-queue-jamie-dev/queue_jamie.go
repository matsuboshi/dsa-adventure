package main 
import "fmt"

// STRUCT THAT REPRESENTS THE QUEUE	
type Queue struct {
	items []int
}

// ENQUEUE ADDS A VALUE AT THE END
func(q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// DEQUEUE REMOVES A VALUE FROM BEGINNING
func (q *Queue) Dequeue() int {
	removed := q.items[0]
	q.items = q.items[1:]
	return removed
}



func main() {
	myQueue := &Queue{}
	myQueue.Enqueue(4)
	myQueue.Enqueue(7)
	myQueue.Enqueue(9)
	myQueue.Enqueue(1)

	fmt.Println(myQueue)

	fmt.Println(myQueue.Dequeue())
	fmt.Println(myQueue.Dequeue())

	fmt.Println(myQueue)
}



