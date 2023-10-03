package main

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}
type MyList[T any] struct {
	head *List[T]
	tail *List[T]
}

func (m *MyList[T]) CreateList() {
	m.head.next = m.tail
}
func main() {
	List := &Mylist[T]
	List.CreateList()
}
