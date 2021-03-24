package lists

type stack struct {
	list *singlyLinkedList
}

func NewStack() *stack {
	return &stack{
		list: NewSinglyLinkedList(nil),
	}
}

func (st *stack) Push(data interface{}) {
	if data != nil {
		st.list.Prepend(data)
	}
}

func (st *stack) Pop() interface{} {
	result := st.list.RemoveHead()
	return result
}

func (st *stack) Peek() interface{}  {
	return st.list.head.data
}

func (st *stack) Size() int {
	return st.list.size
}

func (st *stack) ToSlice() []interface{} {
	return st.list.ToSlice()
}

func (st *stack) PrintAll() {
	st.list.PrintAll()
}
