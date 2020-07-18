package stack

const MAX = 100000

type stack struct{
	items [] interface{}
	top int
}

func Stack(itemLen int) *stack {
	len := itemLen

	if MAX < len{
		len = MAX
	}

	return &stack{
		items: make([]interface{}, len),
		top: 0,
	}
}

func (s *stack) Push (val interface{}) {
	s.items[s.top] = val
	s.top++
}

func (s *stack) Pop () interface{} {
	if s.Empty() == 1 {
		return -1
	}

	s.top--
	return s.items[s.top]
}

func (s *stack) Empty () int {
	if s.top == 0 {
		return 1
	}
	return 0
}

func (s *stack) Top () interface{} {
	if s.Empty() == 1 {
		return -1
	}
	return s.items[s.top - 1]
}

func (s *stack) Size () int {
	return s.top
}
