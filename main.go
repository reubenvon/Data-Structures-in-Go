package main

import (
	"datastructures/src/lists"
	"fmt"
)

func main() {
	str := "Reverse me!"
	stack := lists.NewStack()
	var result []byte
	for _, b := range []byte(str) {
		stack.Push(b)
	}
	stackSize := stack.Size()
	for i := 0; i < stackSize; i++ {
		result = append(result, stack.Pop().(byte))
	}
	fmt.Println(string(result))
}
