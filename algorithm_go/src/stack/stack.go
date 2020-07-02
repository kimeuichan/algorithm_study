package stack

const MAX = 10005

type stack struct{
	items [] int
	top int
}

func Stack(itemLen int) *stack {
	len := itemLen

	if MAX < len{
		len = MAX
	}

	return &stack{
		items: make([]int, len),
		top: 0,
	}
}

func (s *stack) Push (val int) {
	s.items[s.top] = val
	s.top++
}

func (s *stack) Pop () int {
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

func (s *stack) Top () int {
	if s.Empty() == 1 {
		return -1
	}
	return s.items[s.top - 1]
}

func (s *stack) Size () int {
	return s.top
}
