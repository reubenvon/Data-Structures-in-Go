package lists

import "fmt"

type node struct {
	data interface{}
	next *node
}

// SinglyLinkedList ADT
type SinglyLinkedList interface {
	Size() int
	IsEmpty() bool
	Append(interface{})
	Prepend(interface{})
	RemoveHead() interface{}
	Head() interface{}
	Tail() interface{}
}

type singlyLinkedList struct {
	head *node
	size int
}

func NewSinglyLinkedList() *singlyLinkedList {
	return &singlyLinkedList{
		head: nil,
		size: 0,
	}
}

func (sll *singlyLinkedList) Size() int {
	return sll.size
}

func (sll *singlyLinkedList) IsEmpty() bool {
	if sll.head == nil {
		return true
	}
	return false
}

func (sll *singlyLinkedList) Append(data interface{}) {
	if sll.head == nil {
		sll.head = &node{data: data}
		sll.size++
		return
	}
	current := sll.head
	for current != nil {
		if current.next == nil {
			current.next = &node{data: data}
			sll.size++
			return
		} else {
			current = current.next
		}
	}
}

func (sll *singlyLinkedList) Prepend(data interface{}) {
	if sll.head == nil {
		sll.head = &node{data: data}
		sll.size++
		return
	}
	temp := sll.head
	sll.head = &node{data: data}
	sll.head.next = temp
	sll.size++
}

func (sll *singlyLinkedList) RemoveHead() interface{} {
	if sll.head == nil {
		return nil
	}
	data := sll.head.data
	next := sll.head.next
	sll.head = next
	sll.size--
	return data
}

func (sll *singlyLinkedList) Head() interface{} {
	return sll.head.data
}

func (sll *singlyLinkedList) Tail() interface{} {
	if sll.head == nil {
		return nil
	}
	current := sll.head
	for current != nil {
		if current.next == nil {
			break
		}
		current = current.next
	}
	return current.data
}

func (sll *singlyLinkedList) RemoveTail() interface{} {
	if sll.head == nil {
		return nil
	}
	var data interface{}
	var prev *node
	current := sll.head
	for current != nil {
		if current.next == nil {
			if prev == nil {
				sll.head = nil
			} else {
				prev.next = nil
			}
			data = current.data
			break
		}
		prev = current
		current = current.next
	}
	sll.size--
	return data
}

func (sll *singlyLinkedList) Reverse() *singlyLinkedList {
	if sll.head != nil {
		var prev *node = nil
		current := sll.head
		for current != nil {
			next := current.next
			current.next = prev
			prev = current
			current = next
		}
		return &singlyLinkedList{head: prev}
	}
	return &singlyLinkedList{}
}

func (sll *singlyLinkedList) PrintAll()  {
	if sll.head == nil {
		fmt.Println("List is empty.")
		return
	}
	current := sll.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func (sll *singlyLinkedList) ToSlice() []interface{} {
	var result []interface{}
	current := sll.head
	for current != nil {
		result = append(result, current.data)
		current = current.next
	}
	return result
}
