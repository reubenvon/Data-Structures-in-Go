package lists

// Queue ADT
type Queue interface {
	Enqueue(interface{})
	Dequeue() interface{}
	Front() interface{}
	Size() int
	IsEmpty() bool
}

type queue struct {
	list *singlyLinkedList
}

func NewQueue() *queue {
	return &queue{
		list: NewSinglyLinkedList(),
	}
}

func (q *queue) Enqueue(data interface{}) {
	q.list.Prepend(data)
}

func (q *queue) Dequeue() interface{} {
	result := q.list.RemoveTail()
	return result
}

func (q *queue) Front() interface{} {
	return q.list.Tail()
}

func (q *queue) Size() int {
	return q.list.Size()
}

func (q *queue) IsEmpty() bool {
	return q.list.IsEmpty()
}
