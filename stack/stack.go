package stack

// Stack implements a LIFO stack
type Stack []interface{}

//Push an item onto the stack
func (s *Stack) Push(n interface{}) { *s = append(*s, n) }

//Pop an item off of the stack LIFO order
func (s *Stack) Pop() (n interface{}) {
	v := *s
	v, n = v[:len(v)-1], v[len(v)-1]
	*s = v
	return n
}

//Len is the length of th stack
func (s *Stack) Len() int { return len(*s) }

//Clear builds a new stack and replaces the current one
//which will allow the gc to clean up the old interface{} memory
func (s *Stack) Reset() {
	*s = Stack{}
}
