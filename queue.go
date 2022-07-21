package collection

// Node contains information and next node pointer.
type Node struct {
	data interface{}
	next *Node
}

// Queue as linked list data structure.
type Queue struct {
	head *Node
	tail *Node
	size int
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		size: 0,
	}
}

// Add to queue.
func (q *Queue) Add(data interface{}) {
	node := new(Node)
	node.data = data

	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}

	q.size++
}

// Remove from queue. Return nil if queue is empty.
func (q *Queue) Remove() interface{} {
	if q.IsEmpty() {
		return nil
	}

	data := q.head.data
	q.head = q.head.next
	q.size--

	return data
}

// Peek at head of queue without removing. Return nil if queue is empty.
func (q *Queue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.head.data
}

// Return number of elements in queue.
func (q *Queue) Size() int {
	return q.size
}

// Return if queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}
