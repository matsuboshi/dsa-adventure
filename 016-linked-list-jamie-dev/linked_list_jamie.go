package main 

import(
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}

func (l *linkedList) prepend (n *node) {
	second := l.head
	l.head = n 
	l.head.next = second
	l.length++
}

func (l linkedList) printListData(){
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")

}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		println("-> empty list")
		return
	}

	if l.head.data == value{
		l.head = l.head.next
		l.length--
		println("-> head deleted")
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			println("-> not found")
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}



func main() {
	mylist := &linkedList{}
	mylist.prepend(&node{data: 62})
	mylist.prepend(&node{data: 55})
	mylist.prepend(&node{data: 16})
	mylist.prepend(&node{data: 18})
	mylist.prepend(&node{data: 14})
	mylist.prepend(&node{data: 78})
	mylist.prepend(&node{data: 88})
	mylist.printListData()
	mylist.deleteWithValue(18)
	mylist.printListData()
	mylist.deleteWithValue(55)
	mylist.printListData()
	mylist.deleteWithValue(55)
	mylist.printListData()
	mylist.deleteWithValue(88)
	mylist.printListData()
	fmt.Println("LENGTH:", mylist.length)
	
	emptylist := &linkedList{}
	emptylist.deleteWithValue(10)
	fmt.Println("LENGTH:", emptylist.length)
}
