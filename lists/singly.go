package lists

import "fmt"

type node struct {
	Data interface{}
	Next *node
}

type singlyLinkedList struct {
	Head *node
}

func NewSinglyLinkedList(data interface{}) *singlyLinkedList {
	if data != nil {
		return &singlyLinkedList{
			Head: &node{Data: data},
		}
	}
	return &singlyLinkedList{}
}

func (sll *singlyLinkedList) Append(data interface{}) {
	if sll.Head == nil {
		sll.Head = &node{Data: data}
		return
	}
	current := sll.Head
	for current != nil {
		if current.Next == nil {
			current.Next = &node{Data: data}
			return
		} else {
			current = current.Next
		}
	}
}

func (sll *singlyLinkedList) Prepend(data interface{}) {
	if sll.Head == nil {
		sll.Head = &node{Data: data}
		return
	}
	temp := sll.Head
	sll.Head = &node{Data: data}
	sll.Head.Next = temp
}

func (sll *singlyLinkedList) RemoveHead() interface{} {
	if sll.Head == nil {
		return nil
	}
	data := sll.Head.Data
	next := sll.Head.Next
	sll.Head = next
	return data
}

func (sll *singlyLinkedList) Reverse() *singlyLinkedList {
	if sll.Head != nil {
		var prev *node = nil
		current := sll.Head
		for current != nil {
			next := current.Next
			current.Next = prev
			prev = current
			current = next
		}
		return &singlyLinkedList{Head: prev}
	}
	return &singlyLinkedList{}
}

func (sll *singlyLinkedList) PrintAll()  {
	if sll.Head == nil {
		fmt.Println("List is empty.")
		return
	}
	current := sll.Head
	for current != nil {
		fmt.Println(current.Data)
		current = current.Next
	}
}

func (sll *singlyLinkedList) ToSlice() []interface{} {
	var result []interface{}
	current := sll.Head
	for current != nil {
		result = append(result, current.Data)
		current = current.Next
	}
	return result
}
