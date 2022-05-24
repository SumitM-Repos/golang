package stack

type Stack []string

func (s *Stack) is_Empty() bool {
	return len(*s) == 0
}

func (s *Stack) push(val string) {
	*s = append(*s, val)
}

func (s *Stack) pop() string {

	if s.is_Empty() {
		return "Stack is empty"
	}
	last_index := len(*s) - 1
	val := (*s)[last_index]

	*s = (*s)[:last_index]
	return val
}
