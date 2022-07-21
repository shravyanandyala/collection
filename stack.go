package collection

// Stack as linked list data structure.
type Stack struct {
	head *Node
	next *Node
	size int
}

func NewStack() *Stack {
	return &Stack{
		head: nil,
		next: nil,
		size: 0,
	}
}

// Push to stack.
func (s *Stack) Push(data interface{}) {
	node := new(Node)
	node.data = data

	oldHead := s.head
	s.head = node
	node.next = oldHead

	s.size++
}

// Pop from stack. Return nil if stack is empty.
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	data := s.head.data
	s.head = s.head.next
	s.size--

	return data
}

// Peek at top of stack without removing. Return nil if stack is empty.
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}

	return s.head.data
}

// Return number of elements in stack.
func (s *Stack) Size() int {
	return s.size
}

// Return if stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.Size() == 0
}
