package deque

const MAX = 100005

type deque struct {
	items []int
	head int
	tail int
}

func Deque(n int) *deque{
	return &deque{
		items: make([]int, n),
		head:   n / 2,
		tail:  n / 2,
	}
}

func (q *deque) PushFront(val int) {
	q.head--
	q.items[q.head] = val
}

func (q *deque) PushBack(val int) {
	q.items[q.tail] = val
	q.tail++
}

func (q *deque) PopFront() int {
	if q.head == q.tail{
		return -1
	}

	val := q.items[q.head]
	q.head++

	return val
}

func (q *deque) PopBack() int {
	if q.head == q.tail{
		return -1
	}

	q.tail--
	val := q.items[q.tail]

	return val
}

func (q *deque) Front() int {
	if q.Empty() == 1{
		return -1
	}

	return q.items[q.head]
}

func (q *deque) Back() int {
	if q.Empty() == 1{
		return -1
	}

	return q.items[q.tail - 1]
}

func (q *deque) Size() int {
	return q.tail - q.head
}

func (q *deque) Empty() int {
	if q.head == q.tail{
		return 1
	}

	return 0
}


