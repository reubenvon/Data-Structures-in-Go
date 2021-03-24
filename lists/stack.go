package lists

type stack struct {
	list *singlyLinkedList
	size int
}

func NewStack() *stack {
	return &stack{
		list: NewSinglyLinkedList(nil),
		size: 0,
	}
}

func (st *stack) Push(data interface{}) {
	if data != nil {
		st.list.Prepend(data)
		st.size++
	}
}

func (st *stack) Pop() interface{} {
	result := st.list.RemoveHead()
	if result != nil {
		st.size--
	}
	return result
}

func (st *stack) Peek() interface{}  {
	return st.list.Head.Data
}

func (st *stack) Size() int {
	return st.size
}

func (st *stack) ToSlice() []interface{} {
	return st.list.ToSlice()
}

func (st *stack) PrintAll() {
	st.list.PrintAll()
}
