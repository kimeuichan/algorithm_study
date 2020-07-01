package queue

const MAX = 100005

type queue struct {
	items []int
	head int
	tail int
}

func Queue() *queue{
	return &queue{
		items: make([]int, MAX),
		head:  0,
		tail:  0,
	}
}

func (q *queue) Push(val int) {
	q.items[q.tail] = val
	q.tail++
}

func (q *queue) Pop() int {
	if q.head == q.tail{
		return -1
	}

	val := q.items[q.head]
	q.head++

	return val
}

func (q *queue) Front() int {
	if q.Empty() == 1{
		return -1
	}

	return q.items[q.head]
}

func (q *queue) Back() int {
	if q.Empty() == 1{
		return -1
	}

	return q.items[q.tail - 1]
}

func (q *queue) Size() int {
	return q.tail - q.head
}

func (q *queue) Empty() int {
	if q.head == q.tail{
		return 1
	}

	return 0
}

