package utils

type Stack []string

func (s *Stack) Push(value string) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}
	i := len(*s) - 1
	value := (*s)[i]
	*s = (*s)[:i]
	return value, true
}

func (s *Stack) Clear() {
	*s = (*s)[:0]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() (string, bool) {
	if len(*s) == 0 {
		return "", false
	}
	return (*s)[len(*s)-1], true
}
